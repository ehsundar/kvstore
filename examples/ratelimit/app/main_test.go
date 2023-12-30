package main

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/ehsundar/kvstore"
	"github.com/ehsundar/kvstore/examples/ratelimit"
)

type RateLimitTestSuite struct {
	suite.Suite

	r *redis.Client
}

func TestRateLimitTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimitTestSuite))
}

func (s *RateLimitTestSuite) SetupSuite() {
	r := redis.NewClient(&redis.Options{
		Network:    "tcp",
		Addr:       "localhost:6379",
		ClientName: "tests",
	})

	_, err := r.FlushAll(context.Background()).Result()
	s.Nil(err)

	s.r = r
}

func (s *RateLimitTestSuite) TearDownTest() {
	_, err := s.r.FlushAll(context.Background()).Result()
	s.Nil(err)
}

func (s *RateLimitTestSuite) TestShouldReadWriteToStorageRespectingTTL() {
	ctx := context.Background()
	storage := ratelimit.NewRateLimitStore(s.r)

	bucketSizeSeconds := int64(1)
	timeBucket := time.Now().Unix() / bucketSizeSeconds

	for i := int64(1); i <= 10; i++ {
		// you need to calculate bucket whenever a request arrives.
		// we comment it here to avoid unwanted situations and test result inconsistencies
		//timeBucket = time.Now().Unix() / bucketSizeSeconds

		rate, err := storage.Incr(ctx, &ratelimit.CallInfo{
			PathName:   "some/path",
			CallerId:   "caller-1",
			TimeBucket: timeBucket,
		}, 1, kvstore.WithIncrTTL(time.Duration(bucketSizeSeconds)*time.Second, true))
		s.Nil(err)
		s.Equal(i, rate)
	}

	value, err := storage.Get(ctx, &ratelimit.CallInfo{
		PathName:   "some/path",
		CallerId:   "caller-1",
		TimeBucket: timeBucket,
	})
	s.Nil(err)
	s.Equal(int64(10), value)
}
