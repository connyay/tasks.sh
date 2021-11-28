package main

import (
	"fmt"
	"log"

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

	globals := map[string]interface{}{
		"printf": fmt.Printf,
		"logf":   log.Printf,
		"panic":  log.Panicf,
		"dump":   spew.Dump,
	}

	_, err := eval(cli.Script, globals, lib.Loader)
	ctx.FatalIfErrorf(err, "eval")
}

func eval(filename string, globals map[string]interface{}, load starlight.LoadFunc) (map[string]interface{}, error) {
	thread := &starlark.Thread{
		Load: load,
	}
	dict, err := convert.MakeStringDict(globals)
	if err != nil {
		return nil, err
	}
	dict, err = starlark.ExecFile(thread, filename, nil, dict)
	if err != nil {
		return nil, err
	}
	return convert.FromStringDict(dict), nil
}
