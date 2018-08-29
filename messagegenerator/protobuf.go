package messagegenerator

import (
	osc2 "github.com/glynternet/oscli/pkg/osc"
	"github.com/golang/protobuf/proto"
	"github.com/hypebeast/go-osc/osc"
	"github.com/pkg/errors"
)

type protobuf string

// NewProtobuf produces a osc message generator configured at the given address.
// A NewProtobuf generator will produce an osc message with a single argument,
// the bytes marshalled from a proto.Message
func NewProtobuf(address string) (*protobuf, error) {
	a, err := osc2.CleanAddress(address)
	return (*protobuf)(&a), errors.Wrap(err, "cleaning address")
}

// NewProtobuf will produce an osc message with a single argument, the bytes
// marshalled from a proto.Message
func (pb protobuf) GenerateMessage(message proto.Message) (*osc.Message, error) {
	bs, err := proto.Marshal(message)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling proto")
	}
	return osc.NewMessage(string(pb), bs), nil
}
