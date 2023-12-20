package optparse

import (
	"github.com/ehsundar/kvstore/protobuf/kvstore"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func inferKeyOptions(o *kvstore.KVStoreKeyOptions, msg *protogen.Message) {
	if o.Name == "" {
		o.Name = extractKeyName(string(msg.Desc.Name()))
	}

	if o.Mode == nil {
		if len(msg.Fields) == 0 {
			o.Mode = &kvstore.KVStoreKeyOptions_StaticKey{
				StaticKey: &kvstore.KVStoreKeyOptions_StaticKeyMode{
					Key: extractKeyName(string(msg.Desc.Name())),
				},
			}
		} else {
			o.Mode = &kvstore.KVStoreKeyOptions_DynamicKey{
				DynamicKey: &kvstore.KVStoreKeyOptions_DynamicKeyMode{},
			}
		}
	}

	switch v := o.Mode.(type) {
	case *kvstore.KVStoreKeyOptions_StaticKey:
		if v.StaticKey.Key == "" {
			v.StaticKey.Key = extractKeyName(string(msg.Desc.Name()))
		}
	}
}

func inferValueOptions(o *kvstore.KVStoreValueOptions, msg *protogen.Message) {
	if o.Name == "" {
		o.Name = extractValueName(string(msg.Desc.Name()))
	}

	if o.Mode == nil && len(msg.Fields) == 1 {
		f := msg.Desc.Fields().Get(0)
		switch f.Kind() {
		case protoreflect.Int32Kind, protoreflect.Int64Kind,
			protoreflect.Uint32Kind, protoreflect.Uint64Kind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:

			o.Mode = &kvstore.KVStoreValueOptions_Numeral{
				Numeral: &kvstore.KVStoreValueOptions_NumeralValueMode{},
			}
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			o.Mode = &kvstore.KVStoreValueOptions_Numeral{
				Numeral: &kvstore.KVStoreValueOptions_NumeralValueMode{},
			}
		}
	}

	if o.Mode == nil {
		o.Mode = &kvstore.KVStoreValueOptions_General{
			General: &kvstore.KVStoreValueOptions_GeneralValueMode{},
		}
	}
}
