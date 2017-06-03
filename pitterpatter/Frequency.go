package pitterpatter

import (
	"encoding/json"
	"log"
)

type Frequency struct {
	Min *normalisedFloat `json:",omitempty"`
	Max *normalisedFloat `json:",omitempty"`
}

func (f Frequency) String() string {
	json, err := json.Marshal(f)
	if err != nil {
		log.Printf("Error creating Fill json for String method: %s", err.Error())
	}
	return string(json)
}
