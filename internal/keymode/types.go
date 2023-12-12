package keymode

import (
	"errors"

	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
)

var (
	ErrKeyModeNotSupported = errors.New("key mode not supported")
	ErrInvalidStaticKey    = errors.New("invalid static key")
)

type KeyFormat struct {
	Format   string
	VarNames []string
}

type Generator func(options *kvstoreProto.KVStoreKeyOptions) (KeyFormat, error)
