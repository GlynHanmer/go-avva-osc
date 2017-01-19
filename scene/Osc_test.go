package scene

import (
	// "fmt"
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

func Test_SceneCut_Auto(t *testing.T) {
	sc := NewSceneCut(AUTO)
	actual := sc.Osc()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append(string(AUTO))
	testOSCMessageResults(actual, expected, t)
}

func Test_SceneCut_Manual(t *testing.T) {
	sc := NewSceneCut(MANUAL)
	actual := sc.Osc()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append(string(MANUAL))
	testOSCMessageResults(actual, expected, t)
}

func Test_SceneCue_ValidSceneNumber(t *testing.T) {
	sc := NewSceneCue(3)
	actual := sc.Osc()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cue")
	expected.Append(int32(3))
	testOSCMessageResults(actual, expected, t)
}

func testOSCMessageResults(actual, expected *osc.Message, t *testing.T) {
	// fmt.Printf("actual: %s %T\n", actual, actual)
	// fmt.Printf("actual: %s %T\n", actual.Address, actual.Address)
	// fmt.Printf("expected: %s %T\n", expected, expected)
	// fmt.Printf("expected: %s %T\n", expected.Address, expected.Address)
	if !actual.Equals(expected) {
		t.Fatalf("Generated osc message not as expected\n\tactual: %+v\n\texpected: %+v", actual, expected)
	}
}
