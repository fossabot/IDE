// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDFA = "tmpl/ui/fa.tmpl"

//
// Renders HTML of template
// FA with struct types.Dex
func FA(d types.Dex) string {
	return netbFA(d)
}

// Render template with JSON string as
// data.
func netFA(args ...interface{}) string {

	// Get data from JSON
	var d = netcFA(args...)
	return netbFA(d)

}

// template render function
func netbFA(d types.Dex) string {
	localid := templateIDFA
	name := "FA"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcFA(args ...interface{}) (d types.Dex) {

	if len(args) > 0 {
		jsonData := args[0].(string)
		err := parseJSON(jsonData, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}

	return
}
