package varz

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"
)

func Loader(next starlight.LoadFunc, taskID string) (starlight.LoadFunc, func()) {
	filename := "./data/" + taskID + ".varz.json"
	varz, err := restoreVarz(filename)
	if err != nil {
		panic(err)
	}
	varzMod, err := convert.MakeStringDict(map[string]interface{}{
		"set_varz": func(name string, value interface{}) error {
			varz[name] = value
			return nil
		},
		"get_varz": func(name string) (interface{}, bool) {
			value, ok := varz[name]
			return value, ok
		},
	})
	if err != nil {
		panic(err)
	}
	return func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			if module == "varz" {
				return varzMod, nil
			}
			return next(thread, module)
		}, func() {
			saveVarz(filename, varz)
		}
}

func restoreVarz(filename string) (map[string]interface{}, error) {
	varzJSON, err := os.ReadFile(filename)
	if err != nil {
		return map[string]interface{}{}, nil
	}
	varz := map[string]interface{}{}
	err = json.NewDecoder(bytes.NewReader(varzJSON)).Decode(&varz)
	return varz, err
}

func saveVarz(filename string, varz map[string]interface{}) error {
	varzJSON, err := json.MarshalIndent(varz, "", "  ")
	if err != nil {
		log.Printf("Failed marshalling varz json %v", err)
		return err
	}
	err = os.WriteFile(filename, varzJSON, 0644)
	if err != nil {
		log.Printf("Failed writing varz json %v", err)
		return err
	}
	return nil
}
