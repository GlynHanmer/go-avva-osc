// Package canvas provides data structures and adheres to the AvvaOscMessageBuilder interface
package canvas

import (
	"errors"
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

var CANVAS_ON = [...]string{"on", "1", "true"}
var CANVAS_OFF = [...]string{}

// Canvas is a avva.studio visual canvas
type Canvas bool

// NewCanvas generates a new canvas
func NewCanvas(active bool) Canvas {
	return Canvas(active)
}

func NewCanvasFromString(canvas string) (*Canvas, error) {
	c := new(Canvas)
	switch canvas {
	case "on", "1", "true":
		*c = Canvas(true)
		return c, nil
	case "off", "0", "false":
		return c, nil
	case "":
		return nil, errors.New("no canvas string given")
	}
	return nil, fmt.Errorf("canvas string not supported: %s", canvas)
}

// Generate generates an OSC message which can be sent to the avva.studio visuals system.
// Generate method majes Canvas type adhere to AvvaOscMessageBuilder interface
func (c Canvas) Generate() *osc.Message {
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
