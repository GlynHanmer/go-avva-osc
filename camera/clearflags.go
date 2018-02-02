package camera

import (
	"fmt"
)

// ClearFlag is the Clear Flags setting on the Camera in the avva.studio Unity3D project
type ClearFlags string

var (
	skybox    = ClearFlags("skybox")
	solid     = ClearFlags("solid")
	depthOnly = ClearFlags("depth")
)

func NewClearFlag(clearFlag string) (*ClearFlags, error) {
	cf := ClearFlags(clearFlag)
	switch cf {
	case skybox, solid, depthOnly:
	default:
		return nil, fmt.Errorf("invalid ClearFlags '%s'", clearFlag)
	}
	return &cf, nil
}
