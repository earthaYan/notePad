package apiserver

import (
	"context"
	"fmt"

	pb "github.com/marmotedu/api/proto/apiserver/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/earthaYan/notePad/internal/apiserver/config"
	cachev1 "github.com/earthaYan/notePad/internal/apiserver/controller/v1/cache"
	"github.com/earthaYan/notePad/internal/apiserver/store"
	"github.com/earthaYan/notePad/internal/apiserver/store/mysql"
	genericoptions "github.com/earthaYan/notePad/internal/pkg/options"
	genericapiserver "github.com/earthaYan/notePad/internal/pkg/server"
	"github.com/earthaYan/notePad/pkg/log"
	"github.com/earthaYan/notePad/pkg/shutdown"
	"github.com/earthaYan/notePad/pkg/shutdown/shutdownmanagers/posixsignal"
	"github.com/earthaYan/notePad/pkg/storage"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	redisOptions     *genericoptions.RedisOptions
	gRPCAPIServer    *grpcAPIServer
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

// ExtraConfig defines extra configuration for the iam-apiserver.
type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	ServerCert   genericoptions.GeneratableKeyCert
	mysqlOptions *genericoptions.MySQLOptions
	// etcdOptions      *genericoptions.EtcdOptions
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	extraServer, err := extraConfig.complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		gs:               gs,
		redisOptions:     cfg.RedisOptions,
		genericAPIServer: genericServer,
		gRPCAPIServer:    extraServer,
	}

	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	s.initRedisStore()

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
		}

		s.gRPCAPIServer.Close()
		s.genericAPIServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	go s.gRPCAPIServer.Run()

	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.Run()
}

type completedExtraConfig struct {
	*ExtraConfig
}

// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

// New create a grpcAPIServer instance.
func (c *completedExtraConfig) New() (*grpcAPIServer, error) {
	creds, err := credentials.NewServerTLSFromFile(c.ServerCert.CertKey.CertFile, c.ServerCert.CertKey.KeyFile)
	if err != nil {
		log.Fatalf("Failed to generate credentials %s", err.Error())
	}
	opts := []grpc.ServerOption{grpc.MaxRecvMsgSize(c.MaxMsgSize), grpc.Creds(creds)}
	grpcServer := grpc.NewServer(opts...)

	storeIns, _ := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	// storeIns, _ := etcd.GetEtcdFactoryOr(c.etcdOptions, nil)
	store.SetClient(storeIns)
	cacheIns, err := cachev1.GetCacheInsOr(storeIns)
	if err != nil {
		log.Fatalf("Failed to get cache instance: %s", err.Error())
	}

	pb.RegisterCacheServer(grpcServer, cacheIns)

	reflection.Register(grpcServer)

	return &grpcAPIServer{grpcServer, c.Addr}, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

// nolint: unparam
func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		Addr:         fmt.Sprintf("%s:%d", cfg.GRPCOptions.BindAddress, cfg.GRPCOptions.BindPort),
		MaxMsgSize:   cfg.GRPCOptions.MaxMsgSize,
		ServerCert:   cfg.SecureServing.ServerCert,
		mysqlOptions: cfg.MySQLOptions,
		// etcdOptions:      cfg.EtcdOptions,
	}, nil
}

func (s *apiServer) initRedisStore() {
	ctx, cancel := context.WithCancel(context.Background())
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		cancel()

		return nil
	}))

	config := &storage.Config{
		Host:                  s.redisOptions.Host,
		Port:                  s.redisOptions.Port,
		Addrs:                 s.redisOptions.Addrs,
		MasterName:            s.redisOptions.MasterName,
		Username:              s.redisOptions.Username,
		Password:              s.redisOptions.Password,
		Database:              s.redisOptions.Database,
		MaxIdle:               s.redisOptions.MaxIdle,
		MaxActive:             s.redisOptions.MaxActive,
		Timeout:               s.redisOptions.Timeout,
		EnableCluster:         s.redisOptions.EnableCluster,
		UseSSL:                s.redisOptions.UseSSL,
		SSLInsecureSkipVerify: s.redisOptions.SSLInsecureSkipVerify,
	}

	// try to connect to redis
	go storage.ConnectToRedis(ctx, config)
}
