package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/connyay/tasks-sh/lib"
	"github.com/davecgh/go-spew/spew"
	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

var cli struct {
	Script string `flag:"" short:"s" name:"script" help:"Script to evaluate." type:"path"`
}

func main() {
	ctx := kong.Parse(&cli)

	globals, err := convert.MakeStringDict(map[string]interface{}{
		"printf": fmt.Printf,
		"logf":   log.Printf,
		"panic":  log.Panicf,
		"dump":   spew.Dump,
		"env":    getenv,
	})
	ctx.FatalIfErrorf(err, "converting globals")

	_, err = eval(cli.Script, globals, lib.Loader)
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
