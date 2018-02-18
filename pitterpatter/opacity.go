package pitterpatter

import (
	"encoding/json"
)

type Opacity struct {
	Gain *float64 `json:",omitempty"`
	Lift *float64 `json:",omitempty"`
}

func (o Opacity) String() string {
	json, _ := json.Marshal(o)
	return string(json)
}
