// Package camera provides data structures and adheres to the AvvaOscMessageBuilder interface
package camera

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"encoding/json"
	"log"
	"errors"
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

type Camera struct{
	ClearFlags
}

// Osc generates an OSC message which can be sent to the avva.studio visuals system.
// Osc method makes Camera type adhere to AvvaOscMessageBuilder interface
func (c Camera) Osc() *osc.Message {
	json, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error creating camera json: %s",err.Error())
	}
	return osc.NewMessage("/camera", string(json))
}

// String makes Canvas adhere to the Stringer interface
func (c Camera) String() string {
	return fmt.Sprintf("%s %s", "camera", c.ClearFlags)
}