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
	return osc.NewMessage("/canvas", c.int32())
}

func (c Canvas) int32() int32 {
	if bool(c) {
		return 1
	}
	return 0
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
