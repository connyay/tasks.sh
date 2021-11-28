package lib

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/connyay/tasks-sh/lib/twitter"
	"github.com/connyay/tasks-sh/lib/yfinance"
)

var modules = map[string]starlark.StringDict{
	"yfinance": yfinance.Module,
	"twitter":  twitter.Module,
}

func Loader(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	mod, ok := modules[module]
	if !ok {
		return nil, fmt.Errorf("unknown module %q", module)
	}
	return mod, nil
}
