package scene

// cut is an avva.studio scene cut action
type cut struct {
	Type CutType
}

// NewSceneCut creates a new avva.studio scene cut action for a provided scene cut type
func NewSceneCut(cutType CutType) cut {
	return cut{Type: cutType}
}

// actionName returns the action-name string for us in the scene cut OSC message
func (c cut) actionName() string {
	return CUT_ACTION_NAME
}

// adheres to the Striner interface
func (c cut) String() string {
	return "Stringer interface not yet implemented."
}

// The type of cut that the action defines
type CutType string

// Possible valid cut types
var CutTypes = [...]CutType{AUTO, MANUAL, END_OF_SCENE}

const (
	AUTO            CutType = "auto"
	MANUAL          CutType = "manual"
	END_OF_SCENE    CutType = "end_of_scene"
	CUT_ACTION_NAME string  = "cut"
)
