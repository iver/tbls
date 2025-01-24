package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/invopop/jsonschema"
	"github.com/k1LoW/tbls/schema"
)

func main() {
	if err := _main(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func _main() error {
	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.ToSnake
	s := r.Reflect(&schema.SchemaJSON{})
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
