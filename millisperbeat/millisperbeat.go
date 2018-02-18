// Package millisperbeat provides data structures and adheres to the AvvaOscMessageBuilder interface
package millisperbeat

import (
	"github.com/hypebeast/go-osc/osc"
)

type millisPerBeat float32

func NewMillisPerBeat(millis float32) millisPerBeat {
	return millisPerBeat(millis)
}

func (mpb millisPerBeat) Generate() *osc.Message {
	return osc.NewMessage("/audio", "millisperbeat", float32(mpb))
}

func (mpb millisPerBeat) String() string {
	return "/audio millisperbeat"
}
