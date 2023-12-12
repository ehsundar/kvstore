// Code generated by protoc-gen-go-kvstore. DO NOT EDIT.

package example

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
)

// generated code for FeatureX
// storage interface

type FeatureXKVStore interface {
	Get(context.Context, *StaticKey, ...FeatureXCallOption) (*StaticPrimitiveBoolValue, error)
	Set(context.Context, *StaticKey,
		*StaticPrimitiveBoolValue, ...FeatureXCallOption) (*StaticPrimitiveBoolValue, error)
	Del(context.Context, *StaticKey) error
}

type featureXCallOptionContext struct{}

type FeatureXCallOption func(o *featureXCallOptionContext)

// storage construction

func NewFeatureXStore(r *redis.Client, opts ...featureXOption) FeatureXKVStore {
	oc := featureXOptionContext{}

	for _, opt := range opts {
		opt(&oc)
	}

	return &featureXStorage{
		r:    r,
		opts: oc,
	}
}

type featureXOptionContext struct{}

type featureXOption func(o *featureXOptionContext)

// storage implementation

type featureXStorage struct {
	r    *redis.Client
	opts featureXOptionContext
}

func (s *featureXStorage) Get(
	ctx context.Context, key *StaticKey, opts ...FeatureXCallOption) (*StaticPrimitiveBoolValue, error) {

	k, err := key.marshal()
	if err != nil {
		return nil, err
	}

	v, err := s.r.Get(ctx, k).Result()
	if err != nil {
		return nil, err
	}

	msg := &StaticPrimitiveBoolValue{}
	err = msg.unmarshal(v)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *featureXStorage) Set(ctx context.Context, key *StaticKey,
	value *StaticPrimitiveBoolValue, opts ...FeatureXCallOption) (*StaticPrimitiveBoolValue, error) {

	k, err := key.marshal()
	if err != nil {
		return nil, err
	}

	mv, err := value.marshal()
	if err != nil {
		return nil, err
	}

	v, err := s.r.SetArgs(ctx, k, mv, redis.SetArgs{
		Mode:     "",
		TTL:      0,
		ExpireAt: time.Time{},
		Get:      true,
		KeepTTL:  false,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if v != "" {
		msg := &StaticPrimitiveBoolValue{}
		err = msg.unmarshal(v)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}

	return nil, nil
}

func (s *featureXStorage) Del(ctx context.Context, key *StaticKey) error {

	k, err := key.marshal()
	if err != nil {
		return err
	}

	_, err = s.r.Del(ctx, k).Result()
	return err
}

// message marshallers

func (msg *StaticKey) marshal() (string, error) {
	// TODO
	v := fmt.Sprintf("%s:%s", "feature-x", "feature_x_enable")

	return v, nil
}

func (msg *StaticPrimitiveBoolValue) marshal() (string, error) {
	v, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (msg *StaticPrimitiveBoolValue) unmarshal(value string) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(value), msg)
}
