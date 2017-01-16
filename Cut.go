package scene

type cut struct {
	Type CutType
}

func NewSceneCut(cutType CutType) cut {
	return cut{Type: cutType}
}

type CutType string

const (
	AUTO   CutType = "auto"
	MANUAL CutType = "manual"
)

// const (
// 	NO_SCENECHANGE_ARGUMENTS_GIVEN_MESSAGE string = "No arguments given to parseSceneChangeCLIArguments function."
// 	INVALID_NUM_OF_ARGUMENTS               string = "Invalid number of arguments given to method."
// )

// type SceneAction interface {
// 	Type      SceneActionType
// 	Arguments []interface{}
// }

// func NewSceneAction(SceneActionType SceneActionType) SceneAction {
// 	return SceneChange{Type: cutType}
// }

// scenechange cut manual
// scenechange cut auto
// scenechange cue NUM

// func parseSceneChangeCLIArguments(arguments []string) (SceneChange, error) {
// 	sceneChange := SceneChange{}
// 	var err error = nil
// 	switch len(arguments) {
// 	case 0:
// 		err = errors.New(NO_SCENECHANGE_ARGUMENTS_GIVEN_MESSAGE)
// 	case 2:
// 		// action := arguments[0]
// 		// switch action {
// 		// case "change":
// 		// }
// 	default:
// 		err = errors.New(INVALID_NUM_OF_ARGUMENTS)
// 	}
// 	return sceneChange, err
// }
