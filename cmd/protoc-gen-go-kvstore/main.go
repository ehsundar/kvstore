package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/ehsundar/kvstore/internal/keymode"
	"github.com/ehsundar/kvstore/internal/optparse"
	"github.com/ehsundar/kvstore/internal/valuemode"
)

func main() {
	setupLogger()
	log.Infof("started protoc-gen-go-kvstore plugin")

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
		keyFormat, err := keymode.GetKeyFormat(pair.KeyOptions, pair.KeyDesc)
		if err != nil {
			panic(err)
		}

		valueMode, err := valuemode.GetValueMode(pair.ValueOptions, pair.ValueDesc)
		if err != nil {
			panic(err)
		}

		templateCtx.Pairs[name] = storagePair{
			CodeSafeName: strcase.ToCamel(name),
			KeySpecs: keySpecs{
				Opts:        pair.KeyOptions,
				MessageName: string(pair.KeyDesc.Name()),
				KeyFormat:   keyFormat,
			},
			ValueSpecs: valueSpecs{
				Opts:         pair.ValueOptions,
				MessageName:  string(pair.ValueDesc.Name()),
				NumericInt:   valueMode == valuemode.ValueModeNumericInt,
				NumericFloat: valueMode == valuemode.ValueModeNumericFloat,
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