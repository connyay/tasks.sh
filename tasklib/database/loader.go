package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/sourcegraph/starlight"
	"github.com/sourcegraph/starlight/convert"
	"go.starlark.net/starlark"

	_ "modernc.org/sqlite" // sqlite driver
)

func Loader(next starlight.LoadFunc, taskID string) (starlight.LoadFunc, func()) {
	sqlDB := &sql.DB{}
	var opened bool
	dbMod, err := convert.MakeStringDict(map[string]interface{}{
		"db_migrate": func(migrations []interface{}) error {
			for _, migration := range migrations {
				if _, err := sqlDB.Exec(migration.(string)); err != nil {
					return err
				}
			}
			return nil
		},
		"db_exec":  dbExec(sqlDB),
		"db_query": dbQuery(sqlDB),
	})
	if err != nil {
		panic(err)
	}
	close := func() {
		if !opened {
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed closing db %v", err)
		}
	}
	return func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
		if module == "database" {
			err := os.MkdirAll("./data", 0750)
			if err != nil {
				return nil, err
			}
			s, err := sql.Open("sqlite", "./data/"+taskID+".db")
			if err != nil {
				return nil, err
			}
			*sqlDB = *s
			opened = true
			return dbMod, nil
		}
		return next(thread, module)
	}, close
}

func dbExec(db *sql.DB) func(query string, args ...interface{}) (int64, string) {
	return func(query string, args ...interface{}) (int64, string) {
		res, err := db.Exec(query, args...)
		if err != nil {
			return 0, err.Error()
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return 0, err.Error()
		}
		return rows, ""
	}
}

func dbQuery(db *sql.DB) func(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return func(query string, args ...interface{}) ([]map[string]interface{}, error) {
		list, err := db.Query(query, args...)
		if err != nil {
			return nil, err
		}
		defer list.Close()
		fields, err := list.Columns()
		if err != nil {
			return nil, err
		}
		var rows []map[string]interface{}
		for list.Next() {
			scans := make([]interface{}, len(fields))
			row := make(map[string]interface{})

			for i := range scans {
				scans[i] = &scans[i]
			}
			list.Scan(scans...)
			for i, v := range scans {
				var value = ""
				if v != nil {
					value = fmt.Sprintf("%s", v)
				}
				row[fields[i]] = value
			}
			rows = append(rows, row)
		}
		return rows, nil
	}
}
