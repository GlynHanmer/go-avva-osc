package scene

import (
	"bytes"
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

func Test_CanCreate_SceneCut_Auto(t *testing.T) {
	actual := NewSceneCut(AUTO)
	expected := cut{Type: AUTO}
	testSceneCutResults(actual, expected, t)
}

func Test_CanCreate_SceneCut_Manual(t *testing.T) {
	actual := NewSceneCut(MANUAL)
	expected := cut{Type: MANUAL}
	testSceneCutResults(actual, expected, t)
}

func testSceneCutResults(actual, expected cut, t *testing.T) {
	if actual != expected {
		var buffer bytes.Buffer
		buffer.WriteString("Actual scene cut different to expected.\n")
		buffer.WriteString(fmt.Sprintf("Actual: %+v\n", actual))
		buffer.WriteString(fmt.Sprintf("Expected: %+v\n", expected))
		t.Fatal(buffer.String())
	}
}

func Test_SceneCut_Auto_OscMessage(t *testing.T) {
	sc := NewSceneCut(AUTO)
	actual := sc.Osc()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append("auto")
	testOSCMessageResults(actual, expected, t)
}

func Test_SceneCut_Manual_OscMessage(t *testing.T) {
	sc := NewSceneCut(MANUAL)
	actual := sc.Osc()
	expected := osc.NewMessage("/scenechange")
	expected.Append("cut")
	expected.Append("manual")
	testOSCMessageResults(actual, expected, t)
}

func Test_SceneCutActionName(t *testing.T) {
	cut := NewSceneCut(AUTO)
	actual := cut.actionName()
	expected := CUT_ACTION_NAME
	if actual != expected {
		var buffer bytes.Buffer
		buffer.WriteString("Actual scene cut action name different to expected.\n")
		buffer.WriteString(fmt.Sprintf("Actual: %+v\n", actual))
		buffer.WriteString(fmt.Sprintf("Expected: %+v\n", expected))
		t.Fatal(buffer.String())
	}
}
