package scene

import (
	"github.com/hypebeast/go-osc/osc"
)

// Osc builds a valid avva.studio OSC message for a scene cut
func (c cut) Osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append(c.actionName())
	msg.Append(string(c.Type))
	return msg
}

// Osc builds a valid avva.studio OSC message for a scene cue
func (c cue) Osc() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append(c.actionName())
	msg.Append(int32(c.SceneNumber))
	return msg
}
