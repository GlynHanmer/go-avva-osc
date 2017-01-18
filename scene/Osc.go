package scene

import (
	"github.com/hypebeast/go-osc/osc"
)

func (c cut) osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append("cut")
	msg.Append(c.Type)
	return msg
}

func (c cue) osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append("cue")
	msg.Append(c.SceneNumber)
	return msg
}
