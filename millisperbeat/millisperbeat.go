// Package millisperbeat provides data structures and adheres to the AvvaOscMessageBuilder interface
package millisperbeat

import (
	"github.com/hypebeast/go-osc/osc"
)

type MillisPerBeat struct{}

func NewMillisPerBeat() MillisPerBeat {
	return MillisPerBeat{}
}

func (c MillisPerBeat) Osc() *osc.Message {
	var millisPerBeat float32 = 250
	return osc.NewMessage("/audio", "millisperbeat", millisPerBeat)
}

func (c MillisPerBeat) String() string {
	return "/audio millisperbeat"
}
