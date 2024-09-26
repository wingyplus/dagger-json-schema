// dagger-json-schema is a tool to generate json schema from Dagger module config
// struct.
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dagger/dagger/core/modules"
	"github.com/invopop/jsonschema"
)

func main() {
    r := new(jsonschema.Reflector)
    if err := r.AddGoComments("github.com/dagger", "./dagger/core/modules"); err != nil {
        log.Fatal(err)
    }
    
    s := r.Reflect(&modules.ModuleConfig{})
    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "  ")
    enc.Encode(s)
}
