package canvas

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
)

type Canvas bool

func NewCanvas(active bool) Canvas {
	return Canvas(active)
}

func (c Canvas) Osc() *osc.Message {
	return osc.NewMessage("/canvas", bool(c))
	// Does it need to be done like this?
	// msg := osc.NewMessage("/canvas")
	// msg.Append(bool(c))
	// return msg
}

func (c Canvas) String() string {
	var state string
	if bool(c) {
		state = "on"
	} else {
		state = "off"
	}
	return fmt.Sprintf("%s %s", "canvas", state)
}
