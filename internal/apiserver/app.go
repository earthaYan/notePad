package apiserver

import (
	"github.com/earthaYan/notePad/internal/apiserver/options"
	"github.com/earthaYan/notePad/pkg/app"
	"github.com/marmotedu/log"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("notePad API Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)
	return application
}
func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
