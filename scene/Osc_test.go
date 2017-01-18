package scene

import (
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

func Test_SceneCut_Auto(t *testing.T) {
	sc := NewSceneCut(AUTO)
	actual := sc.osc()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append(AUTO)
	// expected := cut{Type: AUTO}.osc()
	testOSCMessageResults(actual, expected, t)
}

func testOSCMessageResults(actual, expected *osc.Message, t *testing.T) {
	if !actual.Equals(expected) {
		t.Fatalf("Generated osc message not as expected\n\tactual: %+v\n\texpected: %+v", actual, expected)
	}
}
