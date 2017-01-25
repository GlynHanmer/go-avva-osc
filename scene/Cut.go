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

type CutType string

const (
	AUTO            CutType = "auto"
	MANUAL          CutType = "manual"
	CUT_ACTION_NAME string  = "cut"
)
