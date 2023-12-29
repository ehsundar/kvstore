package version

import "fmt"

const Version = "v0.1.4"

func PrintVersion(_ string) error {
	fmt.Printf("protoc-gen-go-kvstore %s\n", Version)

	return nil
}
