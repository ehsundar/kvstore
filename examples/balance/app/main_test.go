package main

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/ehsundar/kvstore"
	"github.com/ehsundar/kvstore/examples/balance"
)

type BalanceTestSuite struct {
	suite.Suite

	r *redis.Client
}

func TestBalanceTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceTestSuite))
}

func (s *BalanceTestSuite) SetupSuite() {
	r := redis.NewClient(&redis.Options{
		Network:    "tcp",
		Addr:       "localhost:6379",
		ClientName: "tests",
	})

	_, err := r.FlushAll(context.Background()).Result()
	s.Nil(err)

	s.r = r
}

func (s *BalanceTestSuite) TearDownTest() {
	_, err := s.r.FlushAll(context.Background()).Result()
	s.Nil(err)
}

func (s *BalanceTestSuite) TestShouldReadWriteToStorageRespectingTTL() {
	ctx := context.Background()
	storage := balance.NewBalanceStore(s.r)

	bal, err := storage.Set(ctx, &balance.BalanceKey{Username: "username"}, 11.11, kvstore.WithRetrieveDisabled())
	s.Nil(err)
	s.InDelta(0, bal, 0.001)

	bal, err = storage.Incr(ctx, &balance.BalanceKey{Username: "username"}, -3.01)
	s.Nil(err)
	s.InDelta(8.1, bal, 0.001)
}
