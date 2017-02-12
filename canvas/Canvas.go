// Package canvas provides data structures and adheres to the AvvaOscMessageBuilder interface
package canvas

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
)

// Canvas is a avva.studio visual canvas
type Canvas bool

// NewCanvas generates a new canvas
func NewCanvas(active bool) Canvas {
	return Canvas(active)
}

// Osc generates an OSC message which can be sent to the avva.studio visuals system.
// Osc method majes Canvas type adhere to AvvaOscMessageBuilder interface
func (c Canvas) Osc() *osc.Message {
	return osc.NewMessage("/canvas", c.int32())
}

// int32 returns the int32 type equivalent of whether the canvas is active or not.
// int32 exists because the avva.studio visuals system OSC receiver does not support booleans in OSC messages
func (c Canvas) int32() int32 {
	if bool(c) {
		return 1
	}
	return 0
}

// String makes Canvas adhere to the Stringer interface
func (c Canvas) String() string {
	var state string
	if bool(c) {
		state = "on"
	} else {
		state = "off"
	}
	return fmt.Sprintf("%s %s", "canvas", state)
}
