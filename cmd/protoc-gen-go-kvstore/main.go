package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/generators"
	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/version"
)

func main() {
	setupLogger()
	log.Infof("started protoc-gen-go-kvstore plugin")

	if len(os.Args) > 1 {
		flag.BoolFunc("version", "print protoc-gen-go-kvstore version", version.PrintVersion)
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
