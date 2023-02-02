package apiserver

import "github.com/earthaYan/notePad/pkg/app"

const commandDesc = `The API server validates and configures data ...`

func NewApp(basename string) *app.App {
	application := app.NewApp("api server", basename)
	return application
}
