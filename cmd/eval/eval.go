package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/google/uuid"
	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"

	"github.com/connyay/tasks-sh/tasklib"
	"github.com/connyay/tasks-sh/tasklib/database"
	"github.com/connyay/tasks-sh/tasklib/environment"
	"github.com/connyay/tasks-sh/tasklib/varz"
)

var cli struct {
	Star       string            `flag:"" name:"star" help:"Star script to evaluate." type:"path"`
	Parameters map[string]string `flag:"" name:"parameters" short:"p" help:"Parameters to pass to script."`
	TaskID     string            `flag:"" name:"task-id" help:"Task ID."`
}

func main() {
	ctx := kong.Parse(&cli)
	taskID := cli.TaskID
	if taskID == "" {
		taskID = fmt.Sprintf("%x", sha256.Sum256([]byte(cli.Star)))
	}
	if taskID == "uuid()" {
		taskID = uuid.NewString()
	}
	loader := tasklib.Loader
	database, close := database.Loader(loader, taskID)
	defer close()
	varz, close := varz.Loader(database, taskID)
	defer close()
	loader = environment.Loader(varz)
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
