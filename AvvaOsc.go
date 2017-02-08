package osc

import (
	"github.com/hypebeast/go-osc/osc"
)

// AvvaOscMessageBuilder is an object that can generate and OSC message which can be sent to the avva.studio visuals system.
type AvvaOscMessageBuilder interface {
	Osc() *osc.Message
}
