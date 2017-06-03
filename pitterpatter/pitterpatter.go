package pitterpatter

import (
	"encoding/json"
	"log"
	"github.com/hypebeast/go-osc/osc"
)

type PitterPatter struct {
	Opacity *Opacity `json:",omitempty"`
	Fill *Fill `json:",omitempty"`
	Frequency *Frequency `json:",omitempty"`
}

func (pp PitterPatter) Osc() *osc.Message {
	json, err := json.Marshal(pp)
	if err != nil {
		log.Printf("Error creating PitterPatter json for osc message: %s", err.Error())
	}
	return osc.NewMessage("/pitterpatter", string(json))
}

func (pp PitterPatter) String() string {
	json, err := json.Marshal(pp)
	if err != nil {
		log.Printf("Error creating PitterPatter json for String method: %s", err.Error())
	}
	return string(json)
}