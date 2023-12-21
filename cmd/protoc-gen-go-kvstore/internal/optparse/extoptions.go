package optparse

import (
	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func ExtractKeyOptions(msg *protogen.Message) *kvstoreProto.KVStoreKeyOptions {
	logrus.Debugf("extracting key options from message: %s", msg.Desc.Name())

	if !proto.HasExtension(msg.Desc.Options(), kvstoreProto.E_KeyOptions) {
		return nil
	}

	ext := proto.GetExtension(msg.Desc.Options(), kvstoreProto.E_KeyOptions)
	logrus.Debugf("message options: %+v", ext)

	o, _ := ext.(*kvstoreProto.KVStoreKeyOptions)

	return o
}

func ExtractValueOptions(msg *protogen.Message) *kvstoreProto.KVStoreValueOptions {
	logrus.Debugf("extracting value options from message: %s", msg.Desc.Name())

	if !proto.HasExtension(msg.Desc.Options(), kvstoreProto.E_ValueOptions) {
		return nil
	}

	ext := proto.GetExtension(msg.Desc.Options(), kvstoreProto.E_ValueOptions)
	logrus.Debugf("message options: %+v", ext)

	o, _ := ext.(*kvstoreProto.KVStoreValueOptions)

	return o
}
