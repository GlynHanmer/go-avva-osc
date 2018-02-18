package scene

import (
	"github.com/hypebeast/go-osc/osc"
)

// Generate builds a valid avva.studio OSC message for a scene cut
func (c cut) Generate() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append("cut")
	msg.Append(string(c.Type))
	return msg
}

// Generate builds a valid avva.studio OSC message for a scene cue
func (c cue) Generate() *osc.Message {
	msg := osc.NewMessage("/scenechange")
	msg.Append("cue")
	msg.Append(int32(c.SceneNumber))
	return msg
}
