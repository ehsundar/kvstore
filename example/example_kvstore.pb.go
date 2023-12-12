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
	Get(context.Context, *StaticKey, ...FeatureXCallOption) (*ValueForStaticKey, error)
	Set(context.Context, *StaticKey,
		*ValueForStaticKey, ...FeatureXCallOption) (*ValueForStaticKey, error)
	Del(context.Context, *StaticKey) error
}

type featureXCallOptionContext struct {
	// common
	ttl  time.Duration
	exAt time.Time

	// set
	mode    string
	get     bool
	keepTTL bool

	// get
	getDel bool
}

type FeatureXCallOption func(o *featureXCallOptionContext)

func WithFeatureXModeNX() FeatureXCallOption {
	return func(o *featureXCallOptionContext) {
		o.mode = "NX"
	}
}

func WithFeatureXModeXX() FeatureXCallOption {
	return func(o *featureXCallOptionContext) {
		o.mode = "XX"
	}
}

func WithFeatureXGetDisabled() FeatureXCallOption {
	return func(o *featureXCallOptionContext) {
		o.get = false
	}
}

func WithFeatureXTTL(ttl time.Duration) FeatureXCallOption {
	return func(o *featureXCallOptionContext) {
		o.ttl = ttl
		o.exAt = time.Time{}
		o.keepTTL = false
		o.getDel = false
	}
}

func WithFeatureXExpireAt(eat time.Time) FeatureXCallOption {
	return func(o *featureXCallOptionContext) {
		o.exAt = eat
		o.ttl = 0
		o.keepTTL = false
		o.getDel = false
	}
}

func WithFeatureXGetDel() FeatureXCallOption {
	return func(o *featureXCallOptionContext) {
		o.getDel = true
		o.ttl = 0
		o.exAt = time.Time{}
	}
}

// storage construction

func NewFeatureXStore(r redis.Cmdable, opts ...featureXOption) FeatureXKVStore {
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
	r    redis.Cmdable
	opts featureXOptionContext
}

func (s *featureXStorage) Get(
	ctx context.Context, key *StaticKey, opts ...FeatureXCallOption) (*ValueForStaticKey, error) {

	o := featureXCallOptionContext{}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return nil, err
	}

	v, err := s.r.Get(ctx, k).Result()
	if err != nil {
		return nil, err
	}

	msg := &ValueForStaticKey{}
	err = msg.unmarshal(v)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *featureXStorage) Set(ctx context.Context, key *StaticKey,
	value *ValueForStaticKey, opts ...FeatureXCallOption) (*ValueForStaticKey, error) {

	o := featureXCallOptionContext{
		get:     true,
		keepTTL: true,
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
		Mode:     o.mode,
		TTL:      o.ttl,
		ExpireAt: o.exAt,
		Get:      o.get,
		KeepTTL:  o.keepTTL,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if v != "" {
		msg := &ValueForStaticKey{}
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

	v := fmt.Sprintf("feature-x:feature_x_enabled")

	return v, nil
}

func (msg *ValueForStaticKey) marshal() (string, error) {
	v, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (msg *ValueForStaticKey) unmarshal(value string) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(value), msg)
}

// generated code for RateLimit
// storage interface

type RateLimitKVStore interface {
	Get(context.Context, *DynamicKey, ...RateLimitCallOption) (*RateLimitCount, error)
	Set(context.Context, *DynamicKey,
		*RateLimitCount, ...RateLimitCallOption) (*RateLimitCount, error)
	Del(context.Context, *DynamicKey) error
}

type rateLimitCallOptionContext struct {
	// common
	ttl  time.Duration
	exAt time.Time

	// set
	mode    string
	get     bool
	keepTTL bool

	// get
	getDel bool
}

type RateLimitCallOption func(o *rateLimitCallOptionContext)

func WithRateLimitModeNX() RateLimitCallOption {
	return func(o *rateLimitCallOptionContext) {
		o.mode = "NX"
	}
}

func WithRateLimitModeXX() RateLimitCallOption {
	return func(o *rateLimitCallOptionContext) {
		o.mode = "XX"
	}
}

func WithRateLimitGetDisabled() RateLimitCallOption {
	return func(o *rateLimitCallOptionContext) {
		o.get = false
	}
}

func WithRateLimitTTL(ttl time.Duration) RateLimitCallOption {
	return func(o *rateLimitCallOptionContext) {
		o.ttl = ttl
		o.exAt = time.Time{}
		o.keepTTL = false
		o.getDel = false
	}
}

func WithRateLimitExpireAt(eat time.Time) RateLimitCallOption {
	return func(o *rateLimitCallOptionContext) {
		o.exAt = eat
		o.ttl = 0
		o.keepTTL = false
		o.getDel = false
	}
}

func WithRateLimitGetDel() RateLimitCallOption {
	return func(o *rateLimitCallOptionContext) {
		o.getDel = true
		o.ttl = 0
		o.exAt = time.Time{}
	}
}

// storage construction

func NewRateLimitStore(r redis.Cmdable, opts ...rateLimitOption) RateLimitKVStore {
	oc := rateLimitOptionContext{}

	for _, opt := range opts {
		opt(&oc)
	}

	return &rateLimitStorage{
		r:    r,
		opts: oc,
	}
}

type rateLimitOptionContext struct{}

type rateLimitOption func(o *rateLimitOptionContext)

// storage implementation

type rateLimitStorage struct {
	r    redis.Cmdable
	opts rateLimitOptionContext
}

func (s *rateLimitStorage) Get(
	ctx context.Context, key *DynamicKey, opts ...RateLimitCallOption) (*RateLimitCount, error) {

	o := rateLimitCallOptionContext{}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return nil, err
	}

	v, err := s.r.Get(ctx, k).Result()
	if err != nil {
		return nil, err
	}

	msg := &RateLimitCount{}
	err = msg.unmarshal(v)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *rateLimitStorage) Set(ctx context.Context, key *DynamicKey,
	value *RateLimitCount, opts ...RateLimitCallOption) (*RateLimitCount, error) {

	o := rateLimitCallOptionContext{
		get:     true,
		keepTTL: true,
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
		Mode:     o.mode,
		TTL:      o.ttl,
		ExpireAt: o.exAt,
		Get:      o.get,
		KeepTTL:  o.keepTTL,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if v != "" {
		msg := &RateLimitCount{}
		err = msg.unmarshal(v)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}

	return nil, nil
}

func (s *rateLimitStorage) Del(ctx context.Context, key *DynamicKey) error {

	k, err := key.marshal()
	if err != nil {
		return err
	}

	_, err = s.r.Del(ctx, k).Result()
	return err
}

// message marshallers

func (msg *DynamicKey) marshal() (string, error) {

	v := fmt.Sprintf("rate-limit:%v:%v:%v", msg.RpcName, msg.CallerId, msg.Bucket)

	return v, nil
}

func (msg *RateLimitCount) marshal() (string, error) {
	v, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (msg *RateLimitCount) unmarshal(value string) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(value), msg)
}
