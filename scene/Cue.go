package scene

// cue is an avva.studio scene cue action
type cue struct {
	SceneNumber uint8
}

// NewSceneCue builds a new scene cue action
func NewSceneCue(sceneNumber uint8) cue {
	return cue{sceneNumber}
}

// actionName returns the action-name string for us in the scene cue OSC message
func (c cue) actionName() string {
	return CUE_ACTION_NAME
}

const (
	CUE_ACTION_NAME       string = "cue"
	NEGATIVE_SCENE_NUMBER string = "Negative scene number, invalid."
)
