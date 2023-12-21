package version

import "fmt"

const Version = "v0.1.3"

func PrintVersion(_ string) error {
	fmt.Printf("protoc-gen-go-kvstore %s\n", Version)

	return nil
}
