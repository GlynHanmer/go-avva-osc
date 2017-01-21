package canvas

import (
	"github.com/hypebeast/go-osc/osc"
)

type canvas bool

func Canvas(active bool) canvas {
	return canvas(active)
}

func (c canvas) Osc() *osc.Message {
	return osc.NewMessage("/canvas", bool(c))
}
