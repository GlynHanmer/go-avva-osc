package scene

import (
	"testing"

	"github.com/hypebeast/go-osc/osc"
	"github.com/magiconair/properties/assert"
)

func Test_SceneCut_Auto(t *testing.T) {
	sc := NewSceneCut(auto)
	actual := sc.Generate()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append(string(auto))
	assert.Equal(t, expected, actual)
}

func Test_SceneCut_Manual(t *testing.T) {
	sc := NewSceneCut(manual)
	actual := sc.Generate()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append(string(manual))
	assert.Equal(t, expected, actual)
}

func Test_SceneCue_ValidSceneNumber(t *testing.T) {
	sc := NewSceneCue(3)
	actual := sc.Generate()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cue")
	expected.Append(int32(3))
	assert.Equal(t, expected, actual)
}
