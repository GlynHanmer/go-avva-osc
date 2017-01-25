package scene

import (
	"github.com/hypebeast/go-osc/osc"
)

func (c cut) Osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append(c.actionName())
	msg.Append(string(c.Type))
	return msg
}

func (c cue) Osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append(c.actionName())
	msg.Append(int32(c.SceneNumber))
	return msg
}
