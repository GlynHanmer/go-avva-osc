package scene

// cue is an avva.studio scene cue action
type cue struct {
	SceneNumber uint8
}

// NewSceneCue builds a new scene cue action
func NewSceneCue(sceneNumber uint8) cue {
	return cue{sceneNumber}
}
