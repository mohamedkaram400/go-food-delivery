package pkg

import (
	"encoding/json"
	"log"
)

func DebugJSON(label string, value any) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		log.Printf("DEBUG %s: failed to format value: %v", label, err)
		return
	}

	log.Printf("DEBUG %s:\n%s", label, string(data))
}