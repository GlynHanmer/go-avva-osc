package scene

type cue struct {
	SceneNumber uint8
}

func NewSceneCue(sceneNumber uint8) cue {
	return cue{sceneNumber}
}

const (
	NEGATIVE_SCENE_NUMBER string = "Negative scene number, invalid."
)
