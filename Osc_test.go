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

// func Test_CanCreate_SceneCut_Manual(t *testing.T) {
// 	actual := NewSceneCut(MANUAL)
// 	expected := cut{Type: MANUAL}
// 	testSceneCutResults(actual, expected, t)
// }

// func testSceneCutResults(actual, expected cut, t *testing.T) {
// 	if actual != expected {
// 		var buffer bytes.Buffer
// 		buffer.WriteString("Actual scene cut different to expected.\n")
// 		buffer.WriteString(fmt.Sprintf("Actual: %+v\n", actual))
// 		buffer.WriteString(fmt.Sprintf("Expected: %+v\n", expected))
// 		t.Fatal(buffer.String())
// 	}
// }

// func Test_SceneCut_Auto_OscMessage(t *testing.T) {
// 	sc := NewSceneCut(AUTO)
// 	actual := sc.osc()
// 	expected := osc.NewMessage("/scenechange")
// 	expected.Append("cut")
// 	expected.Append("auto")
// 	testOSCMessageResults(actual, expected, t)
// }

// func Test_SceneCut_Manual_OscMessage(t *testing.T) {
// 	sc := NewSceneCut(MANUAL)
// 	actual := sc.osc()
// 	expected := osc.NewMessage("/scenechange")
// 	expected.Append("cut")
// 	expected.Append("manual")
// 	testOSCMessageResults(actual, expected, t)
// }

// func testErrors(actual, expected error, t *testing.T) {
// 	if actual != expected {
// 		t.Fatalf("Error not as expected.\n\tActual: %+v\n\tExpected: %+v", actual, expected)
// 	}
// }

// // func failTestIfNoError(err error, t *testing.T) {
// // 	if err == nil {
// // 		t.Fatalf("No error returned for empty scene change cli arguments")
// // 	}
// // }

// // test for valid and invalid first arguments
// // test for valid and invalid second arguments

// func Test_ParseCLIArguments_ReturnsError_ForEmptyArguments(t *testing.T) {
// 	_, err := parseSceneChangeCLIArguments([]string{})
// 	failTestIfNoError(err, t)
// 	testErrorMessages(NO_SCENECHANGE_ARGUMENTS_GIVEN_MESSAGE, err, t)
// }

// func Test_ParseCLIArguments_ReturnsError_ForInvalidNumberOfArguments(t *testing.T) {
// 	args := make([]string, 6)
// 	_, err := parseSceneChangeCLIArguments(args)
// 	failTestIfNoError(err, t)
// 	testErrorMessages(INVALID_NUM_OF_ARGUMENTS, err, t)
// }

// func Test_ParseCLIArguments_ReturnerrsError_ForInvalidSceneChangeAction(t *testing.T) {
// 	invalidTestArg := "INVALID_TEST_ARG"
// 	args := []string{invalidTestArg, ""}
// 	_, err := parseSceneChangeCLIArguments(args)
// 	failTestIfNoError(err, t)
// 	testErrorMessages("Invalid scene action given as first argument: "+invalidTestArg, err, t)
// }

// // func Test_ParseCLIArguments_ReturnsError_ForInvalidNumberOfArguments(t *testing.T) {
// // 	// firstArg := "INCORRECT_FIRST_ARG"
// // 	args := make([6]string)
// // 	_, err := parseSceneChangeCLIArguments(args)
// // 	if err == nil {
// // 		t.Fatalf("No error returned for empty scene change cli arguments")
// // 	}
// // 	if expected := fmt.Sprintf("Invalid first argument given to parseSceneChangeCLIArguments function: %s", firstArg); err.Error() != expected {
// // 		t.Fatalf("Wrong error message given for empty arguments.\n\tActual: %s\n\tExpected: %s", err.Error(), expected)
// // 	}
// // }

// // func Test_ParseCLIArguments_SceneChange_Auto(t *testing.T) {
// // 	expected := NewSceneChange(AUTO)
// // 	cliArguments := []string{"sc", "auto"}
// // 	actual, err := parseSceneChangeCLIArguments(cliArguments)
// // 	testSceneChangeResults(actual, expected, t)
// // }
