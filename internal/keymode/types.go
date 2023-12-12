package keymode

import (
	"errors"

	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
)

var (
	ErrKeyModeNotSupported = errors.New("key mode not supported")
	ErrInvalidStaticKey    = errors.New("invalid static key")
	ErrDynamicWithNoFields = errors.New("usage of dynamic key mode with no fields in message")
)

type KeyFormat struct {
	Format   string
	VarNames []string
}

type Generator func(options *kvstoreProto.KVStoreKeyOptions) (KeyFormat, error)
