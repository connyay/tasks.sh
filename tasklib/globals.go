package tasklib

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/lib/time"
	"go.starlark.net/starlark"
)

var Globals starlark.StringDict

func init() {
	var err error
	Globals, err = convert.MakeStringDict(map[string]interface{}{
		"printf":  fmt.Printf,
		"sprintf": fmt.Sprintf,
		"logf":    log.Printf,
		"panic":   log.Panicf,
		"dump":    spew.Dump,

		"time": time.Module,
	})
	if err != nil {
		panic(err)
	}
}
