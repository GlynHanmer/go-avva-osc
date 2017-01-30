package scene

type cut struct {
	Type CutType
}

func NewSceneCut(cutType CutType) cut {
	return cut{Type: cutType}
}

func (c cut) actionName() string {
	return CUT_ACTION_NAME
}

func (c cut) String() string {
	return "Stringer interface not yet implemented."
}

type CutType string

var CutTypes = [...]CutType{AUTO, MANUAL}

const (
	AUTO            CutType = "auto"
	MANUAL          CutType = "manual"
	CUT_ACTION_NAME string  = "cut"
)
