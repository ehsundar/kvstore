package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {

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

	templateCtx.Pairs["FeatureX"] = storagePair{
		CodeSafeName: "FeatureX",
		KeySpecs: keySpecs{
			Opts:        nil,
			MessageName: "StaticKey",
		},
		ValueSpecs: valueSpecs{
			Opts:        nil,
			MessageName: "StaticPrimitiveBoolValue",
		},
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
