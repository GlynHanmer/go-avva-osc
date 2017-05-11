package camera

import (
	"testing"
	"encoding/json"
)

func Test_createCamera(t *testing.T) {
	cf := SKYBOX
	t.Log(cf)
}

func Test_ClearFlagJson(t *testing.T) {
	json, err := json.Marshal(SKYBOX)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(json))
}

func Test_CameraJson(t *testing.T){
	var testSets = []struct {
		originalCamera Camera
	}{
		{originalCamera:Camera{}},
		{originalCamera:Camera{ClearFlags:SKYBOX}},
		{originalCamera:Camera{ClearFlags:SOLID}},
		{originalCamera:Camera{ClearFlags:DEPTH_ONLY}},
		{originalCamera:Camera{FieldOfView:NewFieldOfView(0)}},
		{originalCamera:Camera{FieldOfView:NewFieldOfView(255)}},
		{originalCamera:Camera{FieldOfView:NewFieldOfView(199)}},
		{originalCamera:Camera{ClearFlags:DEPTH_ONLY,FieldOfView:NewFieldOfView(25)}},
	}
	for _, testSet := range testSets {
		json, err := json.Marshal(testSet.originalCamera)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(json))
	}
}

//todo test that this can then be parsed back into a camera vis unmarshalling json??
func Test_CameraOsc(t *testing.T) {
	cam := Camera{ClearFlags: SKYBOX}
	msg := cam.Osc()
	t.Log(msg)
}