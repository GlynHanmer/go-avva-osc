// Package osc builds OSC messages to be sent to the avva.studio visuals generating service.
// osc holds various types that represent entities in the avva.studio visuals system.
package osc

import (
	"github.com/hypebeast/go-osc/osc"
)

// AvvaOscMessageBuilder is an object that can generate and OSC message which can be sent to the avva.studio visuals system.
type AvvaOscMessageBuilder interface {
	Osc() *osc.Message
}
