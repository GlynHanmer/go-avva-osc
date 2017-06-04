package pitterpatter

import (
	"encoding/json"
	"log"
)

type Opacity struct {
	Gain *float64 `json:",omitempty"`
	Lift *float64 `json:",omitempty"`
}

func (o Opacity) String() string {
	json, err := json.Marshal(o)
	if err != nil {
		log.Printf("Error creating Opacity json for String method: %s", err.Error())
	}
	return string(json)
}