package scene

type cue struct {
	SceneNumber uint8
}

func NewSceneCue(sceneNumber uint8) cue {
	return cue{sceneNumber}
}

func (c cue) actionName() string {
	return CUE_ACTION_NAME
}

const (
	CUE_ACTION_NAME       string = "cue"
	NEGATIVE_SCENE_NUMBER string = "Negative scene number, invalid."
)
