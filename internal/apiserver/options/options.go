package options

import (
	genericoptions "github.com/earthaYan/notePad/internal/pkg/options"
	"github.com/marmotedu/log"
)

// Options runs an iam api server.
// Generic通用
type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	GRPCOptions             *genericoptions.GRPCOptions            `json:"grpc"     mapstructure:"grpc"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	RedisOptions            *genericoptions.RedisOptions           `json:"redis"    mapstructure:"redis"`
	JwtOptions              *genericoptions.JwtOptions             `json:"jwt"      mapstructure:"jwt"`
	Log                     *log.Options                           `json:"log"      mapstructure:"log"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
}

func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		GRPCOptions:             genericoptions.NewGRPCOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		RedisOptions:            genericoptions.NewRedisOptions(),
		JwtOptions:              genericoptions.NewJwtOptions(),
		Log:                     log.NewOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
	}

	return &o
}
