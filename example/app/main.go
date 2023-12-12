package main

import (
	"context"
	"fmt"

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
	})
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
}
