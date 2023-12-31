package generators

import (
	"fmt"
	"go/format"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal"
	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/keymode"
	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/optparse"
	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/valuemode"
)

func GenerateFile(gen *protogen.Plugin, file *protogen.File) {
	templateCtx := kvstoreTemplateContext{
		PackageName: string(file.GoPackageName),
		GenVersion:  internal.Version,
		Pairs:       extractKeyValuePairs(file),
	}

	if len(templateCtx.Pairs) == 0 {
		return
	}

	value, err := Render(templateCtx)
	if err != nil {
		fmt.Printf("%s\n", err)

		return
	}

	fmtValue, err := formatGoCode([]byte(value))
	if err != nil {
		panic(err)
	}

	filename := file.GeneratedFilenamePrefix + "_kvstore.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	_, err = g.Write(fmtValue)

	if err != nil {
		fmt.Printf("%s\n", err)

		return
	}

	return
}

func extractKeyValuePairs(file *protogen.File) map[string]storagePair {
	templatePairs := make(map[string]storagePair)
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

		var valueKind protoreflect.Kind
		if valueMode == valuemode.NumericInt || valueMode == valuemode.NumericFloat {
			valueKind = pair.ValueDesc.Fields().Get(0).Kind()
		}

		templatePairs[name] = storagePair{
			CodeSafeName: strcase.ToCamel(name),
			KeySpecs: keySpecs{
				Opts:        pair.KeyOptions,
				MessageName: string(pair.KeyDesc.Name()),
				KeyFormat:   keyFormat,
			},
			ValueSpecs: valueSpecs{
				Opts:         pair.ValueOptions,
				MessageName:  string(pair.ValueDesc.Name()),
				NumericInt:   valueMode == valuemode.NumericInt,
				NumericFloat: valueMode == valuemode.NumericFloat,
				NumericType:  protoKindToGoType(valueKind),
			},
		}
	}

	return templatePairs
}

func protoKindToGoType(k protoreflect.Kind) string {
	switch k {
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.Int32Kind:
		return "int32"
	case protoreflect.Int64Kind:
		return "int64"
	case protoreflect.Uint32Kind:
		return "uint32"
	case protoreflect.Uint64Kind:
		return "uint64"
	default:
		//nolint:godox
		// TODO: Handle all proto types
		return ""
	}
}

func formatGoCode(code []byte) ([]byte, error) {
	formatted, err := format.Source(code)
	if err != nil {
		return []byte{}, err
	}

	//nolint:gomnd
	options := &imports.Options{Comments: true, TabIndent: true, TabWidth: 8}
	formatted, err = imports.Process("", formatted, options)

	if err != nil {
		return []byte{}, err
	}

	return formatted, nil
}
