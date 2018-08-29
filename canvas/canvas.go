// Package canvas provides data structures and adheres to the AvvaOscMessageBuilder interface
package canvas

import (
	"github.com/glynternet/go-avva-osc/avvaproto"
	osc2 "github.com/glynternet/oscli/pkg/osc"
	"github.com/golang/protobuf/proto"
	"github.com/hypebeast/go-osc/osc"
	"github.com/pkg/errors"
)

type Canvas avvaproto.Canvas

// Generate generates an OSC message which can be sent to the avva.studio visuals system.
// Generate method makes Canvas type adhere to AvvaOscMessageBuilder interface
func (c Canvas) Generate() (*osc.Message, error) {
	return bytesMessageGenerator("/canvas").generateMessage((*avvaproto.Canvas)(&c))
}

type bytesMessageGenerator string

func newBytesMessageGenerator(address string) (*bytesMessageGenerator, error) {
	a, err := osc2.CleanAddress(address)
	return (*bytesMessageGenerator)(&a), errors.Wrap(err, "cleaning address")
}

func (bmg bytesMessageGenerator) generateMessage(message proto.Message) (*osc.Message, error) {
	bs, err := proto.Marshal(message)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling proto")
	}
	return osc.NewMessage(string(bmg), bs), nil
}
