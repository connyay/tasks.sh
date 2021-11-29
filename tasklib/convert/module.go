package convert

import (
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

var Module starlark.StringDict

func init() {
	m, err := convert.MakeStringDict(map[string]interface{}{
		"to_string_map": stringMap,
	})
	if err != nil {
		panic("converting convert module")
	}
	Module = m
}
func stringMap(in map[interface{}]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(in))
	for k, v := range in {
		out[k.(string)] = v
	}
	return out
}
