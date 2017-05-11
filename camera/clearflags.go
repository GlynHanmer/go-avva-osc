package camera

import (
	"errors"
	"fmt"
)

// ClearFlag is the Clear Flags setting on the Camera in the avva.studio Unity3D project
type ClearFlags string

var (
	SKYBOX     ClearFlags = ClearFlags("skybox")
	SOLID      ClearFlags = ClearFlags("solid")
	DEPTH_ONLY ClearFlags = ClearFlags("depth")
)

func NewClearFlag(clearFlag string) (ClearFlags, error) {
	var cf ClearFlags
	var err error
	switch ClearFlags(clearFlag) {
	case SKYBOX:
		cf = SKYBOX
	case SOLID:
		cf = SOLID
	case DEPTH_ONLY:
		cf = DEPTH_ONLY
	default:
		err = errors.New(fmt.Sprintf("No clearFlag exists with the type %s", clearFlag))
	}
	return cf, err
}
