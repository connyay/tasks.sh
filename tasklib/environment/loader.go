package environment

import (
	"os"
	"strings"

	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

func Loader(next starlight.LoadFunc) starlight.LoadFunc {
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
