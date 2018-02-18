// Package osc builds OSC messages to be sent to the avva.studio visuals generating service.
// osc holds various types that represent entities in the avva.studio visuals system.
package osc

import (
	"github.com/hypebeast/go-osc/osc"
)

// MessageGenerator generates an osc.Message
type MessageGenerator interface {
	Generate() *osc.Message
}
