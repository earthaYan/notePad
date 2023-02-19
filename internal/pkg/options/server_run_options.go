package options

import "github.com/spf13/pflag"

type ServerRunOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{}
}

func (s *ServerRunOptions) Validate() []error {
	errors := []error{}
	return errors
}
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Mode, "server.mode", s.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")

	fs.BoolVar(&s.Healthz, "server/healthz", s.Healthz, ""+
		"Add self readiness check and install /healthz router.")
	fs.StringSliceVar(&s.Middlewares, "server..Middlewares", s.Middlewares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")
}
