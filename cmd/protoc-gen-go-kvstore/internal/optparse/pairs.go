package optparse

import (
	"errors"
	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrKeyOptionsDoNotMatchValueOptions = errors.New("key and value options does not match")
)

type KVStorePair struct {
	KeyOptions *kvstoreProto.KVStoreKeyOptions
	KeyDesc    protoreflect.MessageDescriptor

	ValueOptions *kvstoreProto.KVStoreValueOptions
	ValueDesc    protoreflect.MessageDescriptor
}

func ExtractPairs(msgs []*protogen.Message) (map[string]KVStorePair, error) {
	pairs := make(map[string]KVStorePair)
	km := map[string]*kvstoreProto.KVStoreKeyOptions{}
	kmsg := map[string]protoreflect.MessageDescriptor{}
	vm := map[string]*kvstoreProto.KVStoreValueOptions{}
	vmsg := map[string]protoreflect.MessageDescriptor{}

	for _, msg := range msgs {
		ko := ExtractKeyOptions(msg)
		if ko != nil {
			log.Infof("parsed key option for %s", msg.Desc.Name())

			inferKeyOptions(ko, msg)

			km[ko.Name] = ko
			kmsg[ko.Name] = msg.Desc
			continue
		}

		vo := ExtractValueOptions(msg)
		if vo != nil {
			log.Infof("parsed value option for %s", msg.Desc.Name())

			inferValueOptions(vo, msg)

			vm[vo.Name] = vo
			vmsg[vo.Name] = msg.Desc
			continue
		}
	}

	if len(km) != len(vm) {
		return nil, ErrKeyOptionsDoNotMatchValueOptions
	}

	for name := range km {
		_, ok := vm[name]
		if !ok {
			return nil, ErrKeyOptionsDoNotMatchValueOptions
		}

		pairs[name] = KVStorePair{
			KeyOptions:   km[name],
			KeyDesc:      kmsg[name],
			ValueOptions: vm[name],
			ValueDesc:    vmsg[name],
		}
	}

	return pairs, nil
}
