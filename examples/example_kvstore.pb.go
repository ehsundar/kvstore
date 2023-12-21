// Code generated by protoc-gen-go-kvstore v0.1.3. DO NOT EDIT.

package examples

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ehsundar/kvstore"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
)

// generated code for OnlineSessions
// storage interface

type OnlineSessionsKVStore interface {
	Get(context.Context, *OnlineSessionsKey, ...kvstore.GetOption) (int64, error)
	Set(context.Context, *OnlineSessionsKey,
		int64, ...kvstore.SetOption) (int64, error)
	Del(context.Context, *OnlineSessionsKey) error
	Incr(context.Context, *OnlineSessionsKey, int64, ...kvstore.IncrOption) (int64, error)
}

// storage construction

func NewOnlineSessionsStore(r redis.Cmdable, opts ...kvstore.InitOption) OnlineSessionsKVStore {

	oc := kvstore.InitOptionContext{}
	for _, opt := range opts {
		opt(&oc)
	}

	return &onlineSessionsStorage{
		r:    r,
		opts: oc,
	}
}

// storage implementation

type onlineSessionsStorage struct {
	r    redis.Cmdable
	opts kvstore.InitOptionContext
}

func (s *onlineSessionsStorage) Get(
	ctx context.Context, key *OnlineSessionsKey, opts ...kvstore.GetOption) (int64, error) {

	var err error

	o := kvstore.GetOptionContext{}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return 0, err
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
		return 0, err
	}

	return strconv.ParseInt(v, 10, 64)

}

func (s *onlineSessionsStorage) Set(ctx context.Context, key *OnlineSessionsKey,
	value int64, opts ...kvstore.SetOption) (int64, error) {

	o := kvstore.SetOptionContext{
		Get:     true,
		KeepTTL: true,
	}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return 0, err
	}

	mv := fmt.Sprintf("%v", value)
	v, err := s.r.SetArgs(ctx, k, mv, redis.SetArgs{
		Mode:     o.Mode,
		TTL:      o.TTL,
		ExpireAt: o.ExAt,
		Get:      o.Get,
		KeepTTL:  o.KeepTTL,
	}).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}

	if v != "" {

		return strconv.ParseInt(v, 10, 64)

	}

	return 0, nil
}

func (s *onlineSessionsStorage) Del(ctx context.Context, key *OnlineSessionsKey) error {

	k, err := key.marshal()
	if err != nil {
		return err
	}

	_, err = s.r.Del(ctx, k).Result()
	return err
}

func (s *onlineSessionsStorage) Incr(ctx context.Context,
	key *OnlineSessionsKey, by int64, opts ...kvstore.IncrOption) (int64, error) {

	o := kvstore.IncrOptionContext{}
	for _, opt := range opts {
		opt(&o)
	}

	k, err := key.marshal()
	if err != nil {
		return 0, err
	}

	v, err := s.r.IncrBy(ctx, k, by).Result()
	if err != nil {
		return 0, err
	}

	switch {
	case o.TTL != 0:
		if o.KeepTTL {
			_, err = s.r.ExpireNX(ctx, k, o.TTL).Result()
		} else {
			_, err = s.r.Expire(ctx, k, o.TTL).Result()
		}
	case !o.ExAt.IsZero():
		_, err = s.r.ExpireAt(ctx, k, o.ExAt).Result()
	}

	return v, err
}

// message marshallers

func (msg *OnlineSessionsKey) marshal() (string, error) {

	v := fmt.Sprintf("online-sessions:online-sessions")

	return v, nil
}

func (msg *OnlineSessionsValue) marshal() (string, error) {
	v, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (msg *OnlineSessionsValue) unmarshal(value string) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(value), msg)
}

// generated code for RateLimit
// storage interface

type RateLimitKVStore interface {
	Get(context.Context, *DynamicKey, ...kvstore.GetOption) (*RateLimitCount, error)
	Set(context.Context, *DynamicKey,
		*RateLimitCount, ...kvstore.SetOption) (*RateLimitCount, error)
	Del(context.Context, *DynamicKey) error
}

// storage construction

func NewRateLimitStore(r redis.Cmdable, opts ...kvstore.InitOption) RateLimitKVStore {

	oc := kvstore.InitOptionContext{}
	for _, opt := range opts {
		opt(&oc)
	}

	return &rateLimitStorage{
		r:    r,
		opts: oc,
	}
}

// storage implementation

type rateLimitStorage struct {
	r    redis.Cmdable
	opts kvstore.InitOptionContext
}

func (s *rateLimitStorage) Get(
	ctx context.Context, key *DynamicKey, opts ...kvstore.GetOption) (*RateLimitCount, error) {

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

	msg := &RateLimitCount{}
	err = msg.unmarshal(v)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *rateLimitStorage) Set(ctx context.Context, key *DynamicKey,
	value *RateLimitCount, opts ...kvstore.SetOption) (*RateLimitCount, error) {

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

// generated code for Static
// storage interface

type StaticKVStore interface {
	Get(context.Context, *StaticKey, ...kvstore.GetOption) (*StaticValue, error)
	Set(context.Context, *StaticKey,
		*StaticValue, ...kvstore.SetOption) (*StaticValue, error)
	Del(context.Context, *StaticKey) error
}

// storage construction

func NewStaticStore(r redis.Cmdable, opts ...kvstore.InitOption) StaticKVStore {

	oc := kvstore.InitOptionContext{}
	for _, opt := range opts {
		opt(&oc)
	}

	return &staticStorage{
		r:    r,
		opts: oc,
	}
}

// storage implementation

type staticStorage struct {
	r    redis.Cmdable
	opts kvstore.InitOptionContext
}

func (s *staticStorage) Get(
	ctx context.Context, key *StaticKey, opts ...kvstore.GetOption) (*StaticValue, error) {

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

	msg := &StaticValue{}
	err = msg.unmarshal(v)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *staticStorage) Set(ctx context.Context, key *StaticKey,
	value *StaticValue, opts ...kvstore.SetOption) (*StaticValue, error) {

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

	if v != "" {
		msg := &StaticValue{}
		err = msg.unmarshal(v)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}

	return nil, nil
}

func (s *staticStorage) Del(ctx context.Context, key *StaticKey) error {

	k, err := key.marshal()
	if err != nil {
		return err
	}

	_, err = s.r.Del(ctx, k).Result()
	return err
}

// message marshallers

func (msg *StaticKey) marshal() (string, error) {

	v := fmt.Sprintf("static:static")

	return v, nil
}

func (msg *StaticValue) marshal() (string, error) {
	v, err := protojson.MarshalOptions{}.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (msg *StaticValue) unmarshal(value string) error {
	return protojson.UnmarshalOptions{}.Unmarshal([]byte(value), msg)
}