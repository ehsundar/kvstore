package main

import "fmt"

const version = "v0.1.3"

func printVersion(_ string) error {
	fmt.Printf("protoc-gen-go-kvstore %s\n", version)
	return nil
}
