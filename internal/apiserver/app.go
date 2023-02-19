package apiserver

import (
	"github.com/earthaYan/notePad/internal/apiserver/config"
	"github.com/earthaYan/notePad/internal/apiserver/options"
	"github.com/earthaYan/notePad/pkg/app"
	"github.com/earthaYan/notePad/pkg/log"
)

const commandDesc = `详细描述`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("notepad api server", basename, app.WithOptions(opts), app.WithDescription(commandDesc), app.WithDefaultValidArgs(), app.WithRunFunc(run(opts)))
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
