package logger

import (
	"github.com/hypebeast/go-osc/osc"
)

type Logger string

func (e *Logger) Generate() *osc.Message {
	return osc.NewMessage("/logger", "address", string(*e))
}
