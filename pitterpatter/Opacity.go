package pitterpatter

import (
	"encoding/json"
	"log"
)

type Opacity struct {
	Gain *normalisedFloat `json:",omitempty"`
	Lift *normalisedFloat `json:",omitempty"`
}

func (o Opacity) String() string {
	json, err := json.Marshal(o)
	if err != nil {
		log.Printf("Error creating Opacity json for String method: %s", err.Error())
	}
	return string(json)
}