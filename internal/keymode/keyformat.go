package keymode

import (
	"fmt"
	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
)

func GetKeyFormat(o *kvstoreProto.KVStoreKeyOptions) (KeyFormat, error) {
	switch v := o.Mode.(type) {
	case *kvstoreProto.KVStoreKeyOptions_StaticKey:
		if err := validateStaticKey(v.StaticKey.Key); err != nil {
			return KeyFormat{}, err
		}

		return KeyFormat{
			Format:   fmt.Sprintf("%s:%s", o.Name, v.StaticKey.Key),
			VarNames: nil,
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
