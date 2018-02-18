package scene

import "fmt"

// cut is an avva.studio scene cut action
type cut struct {
	Type cutType
}

// NewSceneCut creates a new avva.studio scene cut action for a provided scene cut type
func NewSceneCut(cutType cutType) cut {
	return cut{Type: cutType}
}

// adheres to the Striner interface
func (c cut) String() string {
	return "Stringer interface not yet implemented."
}

// The type of cut that the action defines
type cutType string

func NewCutType(t string) (*cutType, error) {
	ct := cutType(t)
	switch ct {
	case auto, manual, endOfScene:
	default:
		return nil, fmt.Errorf("cut type not supported: %s", t)
	}
	return &ct, nil
}

const (
	auto       cutType = "auto"
	manual     cutType = "manual"
	endOfScene cutType = "end_of_scene"
)
