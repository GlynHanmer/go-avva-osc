package scene

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_CanCreate_SceneCue(t *testing.T) {
	actual := NewSceneCue(5)
	expected := cue{5}
	testSceneCueResults(actual, expected, t)
}

func testSceneCueResults(actual, expected cue, t *testing.T) {
	if actual != expected {
		var buffer bytes.Buffer
		buffer.WriteString("Actual scene cue different to expected.\n")
		buffer.WriteString(fmt.Sprintf("Actual: %+v\n", actual))
		buffer.WriteString(fmt.Sprintf("Expected: %+v\n", expected))
		t.Fatal(buffer.String())
	}
}

func Test_SceneCueActionName(t *testing.T) {
	cue := NewSceneCue(3)
	actual := cue.actionName()
	expected := CUE_ACTION_NAME
	if actual != expected {
		var buffer bytes.Buffer
		buffer.WriteString("Actual scene cue action name different to expected.\n")
		buffer.WriteString(fmt.Sprintf("Actual: %+v\n", actual))
		buffer.WriteString(fmt.Sprintf("Expected: %+v\n", expected))
		t.Fatal(buffer.String())
	}
}
