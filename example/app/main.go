package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ehsundar/kvstore/example"
)

func main() {
	ctx := context.Background()

	r := redis.NewClient(&redis.Options{
		Network:    "tcp",
		Addr:       "localhost:6379",
		ClientName: "example",
		DB:         0,
	})
	_, err := r.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	featureX := example.NewFeatureXStore(r)

	_, err = featureX.Set(ctx, &example.StaticKey{}, &example.ValueForStaticKey{
		Value:  true,
		Phones: []string{"123", "456"},
		Items: &example.ValueForStaticKey_NestedItems{
			Items: []int32{1, 2, 3},
		},
	},
		example.WithFeatureXSetTTL(10*time.Second),
	)
	if err != nil {
		panic(err)
	}

	v, err := featureX.Get(ctx, &example.StaticKey{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", v)

	if err := featureX.Del(ctx, &example.StaticKey{}); err != nil {
		panic(err)
	}

	rateLimit := example.NewRateLimitStore(r)

	_, err = rateLimit.Set(ctx, &example.DynamicKey{
		RpcName:  "GetUser",
		CallerId: "caller-one",
		// every ten seconds is a separate bucket
		Bucket: 1702411920,
	}, &example.RateLimitCount{
		Count: 10,
		Limit: 100,
	},
		example.WithRateLimitSetTTL(10*time.Second),
	)
	if err != nil {
		panic(err)
	}

	rateLimitValue, err := rateLimit.Get(ctx, &example.DynamicKey{
		RpcName:  "GetUser",
		CallerId: "caller-one",
		Bucket:   1702411920,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d/%d\n", rateLimitValue.Count, rateLimitValue.Limit)

	sessions := example.NewOnlineSessionsStore(r)
	current, err := sessions.Incr(ctx, &example.OnlineSessionsKey{}, 2,
		example.WithOnlineSessionsIncrTTL(10*time.Second, true),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", current)
}
