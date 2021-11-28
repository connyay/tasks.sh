package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/connyay/tasks-sh/tasklib"
	"github.com/davecgh/go-spew/spew"
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

	globals, err := convert.MakeStringDict(map[string]interface{}{
		"printf":     fmt.Printf,
		"sprintf":    fmt.Sprintf,
		"logf":       log.Printf,
		"panic":      log.Panicf,
		"dump":       spew.Dump,
		"env":        getenv,
		"parameters": cli.Parameters,
	})
	ctx.FatalIfErrorf(err, "converting globals")

	_, err = eval(cli.Star, globals, tasklib.Loader)
	ctx.FatalIfErrorf(err, "eval")
}

func getenv(k string) string {
	// allowlist? namespaced?
	return os.Getenv(k)
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
