package main

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/ehsundar/kvstore"
	"github.com/ehsundar/kvstore/examples/flags"
)

type FlagsTestSuite struct {
	suite.Suite

	r *redis.Client
}

func TestFlagsTestSuite(t *testing.T) {
	suite.Run(t, new(FlagsTestSuite))
}

func (s *FlagsTestSuite) SetupSuite() {
	r := redis.NewClient(&redis.Options{
		Network:    "tcp",
		Addr:       "localhost:6379",
		ClientName: "tests",
	})

	_, err := r.FlushAll(context.Background()).Result()
	s.Nil(err)

	s.r = r
}

func (s *FlagsTestSuite) TearDownTest() {
	_, err := s.r.FlushAll(context.Background()).Result()
	s.Nil(err)
}

func (s *FlagsTestSuite) TestShouldReadWriteToStorage() {
	ctx := context.Background()
	storage := flags.NewFlagStore(s.r)

	_, err := storage.Set(ctx, &flags.FlagKey{}, &flags.FlagValue{
		Endpoint:          "test_value",
		ExperimentPercent: 24,
	})
	s.Nil(err)

	value, err := storage.Get(ctx, &flags.FlagKey{})
	s.Nil(err)

	s.Equal("test_value", value.Endpoint)
	s.Equal(int32(24), value.ExperimentPercent)
}

func (s *FlagsTestSuite) TestShouldRespectTTL() {
	ctx := context.Background()
	storage := flags.NewFlagStore(s.r)

	_, err := storage.Set(ctx, &flags.FlagKey{}, &flags.FlagValue{}, kvstore.WithSetTTL(100*time.Millisecond))
	s.Nil(err)

	_, err = storage.Get(ctx, &flags.FlagKey{})
	s.Nil(err)

	time.Sleep(100 * time.Millisecond)

	_, err = storage.Get(ctx, &flags.FlagKey{})
	s.ErrorIs(err, redis.Nil)
}
