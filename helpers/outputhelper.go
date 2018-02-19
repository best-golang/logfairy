package helpers

import (
	"encoding/json"
	"log"
)

// JSONPrettyPrint print the pretty json representation of an object
func JSONPrettyPrint(object interface{}) {
	prettyObject, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s", prettyObject)
}
