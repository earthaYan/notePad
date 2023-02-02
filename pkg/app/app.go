package app

import "github.com/spf13/cobra"

type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	commands    []*Command
	args        cobra.PositionalArgs
	cmd         *cobra.Command
}

// type NewInt int —— 将NewInt定义为int类型,属于新的类型
// type IntAlias = int —— 将int取一个别名叫IntAlias
// 不能在 不是 main 包中定义的类型上定义新方法
type Option func(*App)

func NewApp(name, basename string, opts ...Option) *App {
	a := &App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(a)
	}
	a.buildCommand()
	return a
}
