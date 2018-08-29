package messagegenerator_test

import (
	"testing"

	"github.com/glynternet/go-avva-osc/avvaproto"
	"github.com/glynternet/go-avva-osc/messagegenerator"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestNewProtobuf(t *testing.T) {
	t.Run("invalid address", func(t *testing.T) {
		_, err := messagegenerator.NewProtobuf("/oia 98")
		assert.Error(t, err)
	})

	t.Run("valid address", func(t *testing.T) {
		g, err := messagegenerator.NewProtobuf("/dummy")
		assert.NoError(t, err)
		assert.NotNil(t, g)
	})
}

func TestProtobuf_GenerateMessage(t *testing.T) {
	g, err := messagegenerator.NewProtobuf("/dummy")
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	msg, err := g.GenerateMessage(&avvaproto.Canvas{
		Id:      6,
		Enabled: true,
	})
	if !assert.NoError(t, err) {
		assert.Nil(t, msg)
		t.FailNow()
	}

	assert.Len(t, msg.Arguments, 1)
	assert.IsType(t, "", msg.Arguments[0])

	bs := []byte(msg.Arguments[0].(string))
	c := &avvaproto.Canvas{}
	err = proto.Unmarshal(bs, c)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	assert.Equal(t,
		avvaproto.Canvas{Id: 6, Enabled: true},
		*c,
	)
}
