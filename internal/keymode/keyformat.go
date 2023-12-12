package keymode

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
	"google.golang.org/protobuf/reflect/protoreflect"

	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
)

func GetKeyFormat(o *kvstoreProto.KVStoreKeyOptions, d protoreflect.MessageDescriptor) (KeyFormat, error) {
	switch v := o.Mode.(type) {
	case *kvstoreProto.KVStoreKeyOptions_StaticKey:
		if err := validateStaticKey(v.StaticKey.Key); err != nil {
			return KeyFormat{}, err
		}

		return KeyFormat{
			Format:   fmt.Sprintf("%s:%s", o.Name, v.StaticKey.Key),
			VarNames: nil,
		}, nil
	case *kvstoreProto.KVStoreKeyOptions_DynamicKey:
		vars := extractVars(d)
		if len(vars) == 0 {
			return KeyFormat{}, ErrDynamicWithNoFields
		}

		f := strings.Join(lo.Times(len(vars), func(_ int) string {
			return "%v"
		}), ":")
		return KeyFormat{
			Format:   fmt.Sprintf("%s:%s", o.Name, f),
			VarNames: vars,
		}, nil
	default:
		return KeyFormat{}, ErrKeyModeNotSupported
	}
}

func validateStaticKey(k string) error {
	if len(k) == 0 {
		return ErrInvalidStaticKey
	}
	return nil
}

func extractVars(d protoreflect.MessageDescriptor) []string {
	return lo.Times(d.Fields().Len(), func(i int) string {
		return strcase.ToCamel(d.Fields().Get(i).TextName())
	})
}
