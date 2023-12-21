package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ehsundar/kvstore"
	"github.com/ehsundar/kvstore/examples"
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

	featureX := examples.NewStaticStore(r)

	_, err = featureX.Set(ctx, &examples.StaticKey{}, &examples.StaticValue{
		Value:  true,
		Phones: []string{"123", "456"},
		Items: &examples.StaticValue_NestedItems{
			Items: []int32{1, 2, 3},
		},
	},
		kvstore.WithSetTTL(10*time.Second),
	)
	if err != nil {
		panic(err)
	}

	v, err := featureX.Get(ctx, &examples.StaticKey{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", v)

	if err := featureX.Del(ctx, &examples.StaticKey{}); err != nil {
		panic(err)
	}

	rateLimit := examples.NewRateLimitStore(r)

	_, err = rateLimit.Set(ctx, &examples.DynamicKey{
		RpcName:  "GetUser",
		CallerId: "caller-one",
		// every ten seconds is a separate bucket
		Bucket: 1702411920,
	}, &examples.RateLimitCount{
		Count: 10,
		Limit: 100,
	},
		kvstore.WithSetTTL(10*time.Second),
	)
	if err != nil {
		panic(err)
	}

	rateLimitValue, err := rateLimit.Get(ctx, &examples.DynamicKey{
		RpcName:  "GetUser",
		CallerId: "caller-one",
		Bucket:   1702411920,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d/%d\n", rateLimitValue.Count, rateLimitValue.Limit)

	sessions := examples.NewOnlineSessionsStore(r)
	current, err := sessions.Incr(ctx, &examples.OnlineSessionsKey{}, 2,
		kvstore.WithIncrTTL(10*time.Second, true),
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", current)
}
