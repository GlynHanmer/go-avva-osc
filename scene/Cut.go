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
