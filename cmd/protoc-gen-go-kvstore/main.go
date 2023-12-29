package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal"
	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/generators"
)

func main() {
	setupLogger()
	log.Infof("started protoc-gen-go-kvstore plugin")

	if len(os.Args) > 1 {
		flag.BoolFunc("version", "print protoc-gen-go-kvstore version", func(_ string) error {
			fmt.Printf("%s %s %s\n", internal.Version, internal.Commit, internal.Date)
			return nil
		})
		flag.Parse()

		return
	}

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generators.GenerateFile(gen, f)
		}

		return nil
	})
}
