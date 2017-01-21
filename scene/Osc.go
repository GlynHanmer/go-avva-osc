package scene

import (
	"github.com/hypebeast/go-osc/osc"
)

func (c cut) Osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	//TODO make these take c.actionName or something like that.
	msg.Append("cut")
	msg.Append(string(c.Type))
	return msg
}

func (c cue) Osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append("cue")
	msg.Append(int32(c.SceneNumber))
	return msg
}
