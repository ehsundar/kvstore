package main

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"

	"github.com/ehsundar/kvstore"
	"github.com/ehsundar/kvstore/examples/views"
)

type ViewsTestSuite struct {
	suite.Suite

	r *redis.Client
}

func TestViewsTestSuite(t *testing.T) {
	suite.Run(t, new(ViewsTestSuite))
}

func (s *ViewsTestSuite) SetupSuite() {
	r := redis.NewClient(&redis.Options{
		Network:    "tcp",
		Addr:       "localhost:6379",
		ClientName: "tests",
	})

	_, err := r.FlushAll(context.Background()).Result()
	s.Nil(err)

	s.r = r
}

func (s *ViewsTestSuite) TearDownTest() {
	_, err := s.r.FlushAll(context.Background()).Result()
	s.Nil(err)
}

func (s *ViewsTestSuite) TestShouldReadWriteToStorageRespectingTTL() {
	ctx := context.Background()
	storage := views.NewGetItemsStore(s.r)

	key := &views.GetItemsRequest{
		ViewId:  12,
		Filters: []string{"some", "filters"},
	}

	_, err := storage.Set(ctx, key, &views.GetItemsResponse{
		Title: "test_title",
		Items: []*views.Item{
			{
				Id:      1,
				Visible: true,
				Display: "display 1",
			},
			{
				Id:      2,
				Visible: false,
				Display: "display 2",
			},
		},
	}, kvstore.WithSetTTL(100*time.Second))
	s.Nil(err)

	value, err := storage.Get(ctx, key, kvstore.WithGetTTL(time.Millisecond))
	s.Nil(err)

	s.Equal("test_title", value.Title)
	s.Equal(2, len(value.Items))
	s.Equal("display 1", value.Items[0].Display)

	time.Sleep(3 * time.Millisecond)

	_, err = storage.Get(ctx, key)
	s.ErrorIs(err, redis.Nil)
}
