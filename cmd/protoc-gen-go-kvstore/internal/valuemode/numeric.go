package valuemode

import (
	"errors"

	"github.com/ehsundar/kvstore/protobuf/kvstore"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrMsgNoOneNumericField  = errors.New("message does not contain exactly one numeric field")
	ErrValueModeNotSupported = errors.New("value mode not supported")
)

const (
	General = iota
	NumericInt
	NumericFloat
)

type ValueMode int

func GetValueMode(o *kvstore.KVStoreValueOptions, d protoreflect.MessageDescriptor) (ValueMode, error) {
	switch o.Mode.(type) {
	case *kvstore.KVStoreValueOptions_General:
		return General, nil
	case *kvstore.KVStoreValueOptions_Numeral:
		if d.Fields().Len() != 1 {
			return 0, ErrMsgNoOneNumericField
		}

		f := d.Fields().Get(0)

		switch f.Kind() {
		case protoreflect.Int32Kind, protoreflect.Int64Kind,
			protoreflect.Uint32Kind, protoreflect.Uint64Kind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
			return NumericInt, nil
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			return NumericFloat, nil
		default:
			return 0, ErrMsgNoOneNumericField
		}
	default:
		return General, nil
	}
}
