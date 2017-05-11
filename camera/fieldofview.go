package camera

type FieldOfView *uint8

func NewFieldOfView(x uint8) FieldOfView {
	return FieldOfView(&x)
}