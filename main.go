package main

import (
	"fmt"
	"github.com/ehsundar/kvstore/internal/codesafe"
	"github.com/ehsundar/kvstore/internal/optparse"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	setupLogger()
	log.Infof("started protoc-gen-kvstore plugin")

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {

	templateCtx := kvstoreTemplateContext{
		PackageName: string(file.GoPackageName),
		Pairs:       map[string]storagePair{},
	}

	pairs, err := optparse.ExtractPairs(file.Messages)
	if err != nil {
		panic(err)
	}

	for name, pair := range pairs {
		templateCtx.Pairs[name] = storagePair{
			CodeSafeName: codesafe.ToCamelCase(name),
			KeySpecs: keySpecs{
				Opts:        pair.KeyOptions,
				MessageName: string(pair.KeyDesc.Name()),
			},
			ValueSpecs: valueSpecs{
				Opts:        pair.ValueOptions,
				MessageName: string(pair.ValueDesc.Name()),
			},
		}
	}

	value, err := Render(templateCtx)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	filename := file.GeneratedFilenamePrefix + "_kvstore.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	_, err = g.Write([]byte(value))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	return
}
