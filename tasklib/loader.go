package tasklib

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/connyay/tasks-sh/tasklib/reddit"
	"github.com/connyay/tasks-sh/tasklib/twilio"
	"github.com/connyay/tasks-sh/tasklib/twitter"
	"github.com/connyay/tasks-sh/tasklib/yfinance"
)

var modules = map[string]starlark.StringDict{
	"yfinance": yfinance.Module,
	"twitter":  twitter.Module,
	"reddit":   reddit.Module,
	"twilio":   twilio.Module,
}

func Loader(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	mod, ok := modules[module]
	if !ok {
		return nil, fmt.Errorf("unknown module %q", module)
	}
	return mod, nil
}
