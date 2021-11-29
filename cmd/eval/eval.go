package main

import (
	"github.com/alecthomas/kong"
	"github.com/connyay/tasks-sh/tasklib"
	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

var cli struct {
	Star       string            `flag:"" name:"star" help:"Star script to evaluate." type:"path"`
	Parameters map[string]string `flag:"" name:"parameters" short:"p" help:"Parameters to pass to script."`
}

func main() {
	ctx := kong.Parse(&cli)

	loader := tasklib.Loader
	database, close := tasklib.DatabaseLoader(loader)
	defer close()
	loader = tasklib.EnvLoader(database)
	err := eval(cli.Star, cli.Parameters, tasklib.Globals, loader)
	ctx.FatalIfErrorf(err, "eval")
}

func eval(filename string, args map[string]string, globals starlark.StringDict, load starlight.LoadFunc) error {
	argsVal, err := convert.ToValue(args)
	if err != nil {
		return err
	}
	thread := &starlark.Thread{
		Load: load,
	}
	dict, err := starlark.ExecFile(thread, filename, nil, globals)
	if err != nil {
		return err
	}
	main := dict["main"]
	_, err = starlark.Call(thread, main, starlark.Tuple{argsVal}, nil)
	return err
}
