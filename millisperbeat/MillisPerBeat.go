// millisperbeat provides data structures and complies with the AvvaOscMessageBuilder interface
package millisperbeat

import (
	"github.com/hypebeast/go-osc/osc"
)

type MillisPerBeat struct{}

func NewMillisPerBeat() MillisPerBeat {
	return MillisPerBeat{}
}

func (c MillisPerBeat) Osc() *osc.Message {
	return osc.NewMessage("/millisperbeat")
}

func (c MillisPerBeat) String() string {
	return "millisperbeat"
}
