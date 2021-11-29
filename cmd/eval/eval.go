package main

import (
	"os"
	"strings"

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

	err := eval(cli.Star, cli.Parameters, tasklib.Globals, envLoader(tasklib.Loader))
	ctx.FatalIfErrorf(err, "eval")
}

func envLoader(next starlight.LoadFunc) starlight.LoadFunc {
	env := map[string]interface{}{}
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		// allowlist? namespaced?
		env[parts[0]] = parts[1]
	}
	envMod, err := convert.MakeStringDict(env)
	if err != nil {
		panic(err)
	}
	return func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
		if module == "environment" {
			return envMod, nil
		}
		return next(thread, module)
	}
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
