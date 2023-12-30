// Code generated by protoc-gen-go-kvstore . DO NOT EDIT.

package views

import (
	"context"
	"errors"
	"fmt"

	"github.com/ehsundar/kvstore"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
)

// generated code for GetItems
// storage interface

type GetItemsKVStore interface {
	Get(context.Context, *GetItemsRequest, ...kvstore.GetOption) (*GetItemsResponse, error)
	Set(context.Context, *GetItemsRequest,
		*GetItemsResponse, ...kvstore.SetOption) (*GetItemsResponse, error)
	Del(context.Context, *GetItemsRequest) error
}

// storage construction

func NewGetItemsStore(r redis.Cmdable, opts ...kvstore.InitOption) GetItemsKVStore {

	oc := kvstore.InitOptionContext{}
	for _, opt := range opts {
		opt(&oc)
	}

	return &getItemsStorage{
		r:    r,
		opts: oc,
	}
}

// storage implementation

type getItemsStorage struct {
	r    redis.Cmdable
	opts kvstore.InitOptionContext
}

func (s *getItemsStorage) Get(
	ctx context.Context, key *GetItemsRequest, opts ...kvstore.GetOption) (*GetItemsResponse, error) {

	var err error

	o := kvstore.GetOptionContext{}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return nil, err
	}

	var v string
	switch {
	case o.Del:
		v, err = s.r.GetDel(ctx, k).Result()
	case o.TTL != 0:
		v, err = s.r.GetEx(ctx, k, o.TTL).Result()
	case !o.ExAt.IsZero():
		// TODO: PR to go-redis for exAt
		err = errors.New("exat is not supported by go-redis")
	default:
		v, err = s.r.Get(ctx, k).Result()
	}

	if err != nil {
		return nil, err
	}

	msg := &GetItemsResponse{}
	err = msg.unmarshal(v)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *getItemsStorage) Set(ctx context.Context, key *GetItemsRequest,
	value *GetItemsResponse, opts ...kvstore.SetOption) (*GetItemsResponse, error) {

	o := kvstore.SetOptionContext{
		Get:     true,
		KeepTTL: true,
	}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return nil, err
	}

	mv, err := value.marshal()
	if err != nil {
		return nil, err
	}

	v, err := s.r.SetArgs(ctx, k, mv, redis.SetArgs{
		Mode:     o.Mode,
		TTL:      o.TTL,
		ExpireAt: o.ExAt,
		Get:      o.Get,
		KeepTTL:  o.KeepTTL,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if v != "" && v != "OK" {
		msg := &GetItemsResponse{}
		err = msg.unmarshal(v)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}

	return nil, nil
}

func (s *getItemsStorage) Del(ctx context.Context, key *GetItemsRequest) error {

	k, err := key.marshal()
	if err != nil {
		return err
	}

	_, err = s.r.Del(ctx, k).Result()
	return err
}

// message marshallers

func (msg *GetItemsRequest) marshal() (string, error) {

	v := fmt.Sprintf("get-items:%v:%v", msg.ViewId, msg.Filters)

	return v, nil
}

func (msg *GetItemsResponse) marshal() (string, error) {
	v, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (msg *GetItemsResponse) unmarshal(value string) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(value), msg)
}
