package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/connyay/tasks-sh/tasklib"
	"github.com/davecgh/go-spew/spew"
	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/lib/time"
	"go.starlark.net/starlark"
)

var cli struct {
	Star       string            `flag:"" name:"star" help:"Star script to evaluate." type:"path"`
	Parameters map[string]string `flag:"" name:"parameters" short:"p" help:"Parameters to pass to script."`
}

func main() {
	ctx := kong.Parse(&cli)

	globals, err := convert.MakeStringDict(map[string]interface{}{
		"printf":     fmt.Printf,
		"sprintf":    fmt.Sprintf,
		"logf":       log.Printf,
		"panic":      log.Panicf,
		"dump":       spew.Dump,
		"parameters": cli.Parameters,

		"time": time.Module,
	})
	ctx.FatalIfErrorf(err, "converting globals")

	_, err = eval(cli.Star, globals, envLoader(tasklib.Loader))
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

func eval(filename string, globals starlark.StringDict, load starlight.LoadFunc) (map[string]interface{}, error) {
	thread := &starlark.Thread{
		Load: load,
	}
	dict, err := starlark.ExecFile(thread, filename, nil, globals)
	if err != nil {
		return nil, err
	}
	return convert.FromStringDict(dict), nil
}
