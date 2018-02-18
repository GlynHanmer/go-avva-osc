package camera

import (
	"encoding/json"
	"testing"

	"github.com/glynternet/go-money/common"
	"github.com/magiconair/properties/assert"
)

func Test_createCamera(t *testing.T) {
	cf := skybox
	t.Log(cf)
}

func Test_ClearFlagJson(t *testing.T) {
	json, err := json.Marshal(skybox)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(json))
}

func Test_CameraJson(t *testing.T) {
	for _, test := range []struct {
		camera Camera
	}{
		{camera: Camera{}},
		{camera: Camera{ClearFlags: skybox}},
		{camera: Camera{ClearFlags: solid}},
		{camera: Camera{ClearFlags: depthOnly}},
		{camera: Camera{FieldOfView: NewFieldOfView(0)}},
		{camera: Camera{FieldOfView: NewFieldOfView(255)}},
		{camera: Camera{FieldOfView: NewFieldOfView(199)}},
		{camera: Camera{ClearFlags: depthOnly, FieldOfView: NewFieldOfView(25)}},
	} {
		bs, err := json.Marshal(test.camera)
		common.FatalIfError(t, err, "marshalling camera json")
		var camera Camera
		err = json.Unmarshal(bs, &camera)
		assert.Equal(t, test.camera, camera)
	}
}

//todo test that this can then be parsed back into a camera vis unmarshalling json??
func Test_CameraOsc(t *testing.T) {
	cam := Camera{ClearFlags: skybox}
	msg := cam.Generate()
	t.Log(msg)
}
