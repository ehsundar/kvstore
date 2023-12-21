package generators

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/samber/lo"

	"github.com/ehsundar/kvstore/cmd/protoc-gen-go-kvstore/internal/keymode"
	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
)

//go:embed kvstore.tmpl
var rawTemplate string

var funcs = template.FuncMap{
	"funcCallArgs": func(varName string, v interface{}) string {
		vv, ok := v.([]string)

		if !ok {
			return ""
		}

		vv = lo.Map(vv, func(item string, _ int) string {
			return fmt.Sprintf("%s.%s", varName, item)
		})

		return strings.Join(vv, ", ")
	},
}

type kvstoreTemplateContext struct {
	PackageName string
	GenVersion  string

	Pairs map[string]storagePair
}

type storagePair struct {
	// storage item name turned into safe string
	// example: feature-x -> FeatureX
	CodeSafeName string
	KeySpecs     keySpecs
	ValueSpecs   valueSpecs
}

type keySpecs struct {
	Opts *kvstoreProto.KVStoreKeyOptions

	MessageName string
	KeyFormat   keymode.KeyFormat
}

type valueSpecs struct {
	Opts *kvstoreProto.KVStoreValueOptions

	MessageName  string
	NumericInt   bool
	NumericFloat bool
	NumericType  string
}

func Render(templateContext kvstoreTemplateContext) (string, error) {
	tmpl := template.Must(
		template.New("base").Funcs(sprig.FuncMap()).Funcs(funcs).Parse(rawTemplate),
	)

	builder := &bytes.Buffer{}
	err := tmpl.Execute(builder, templateContext)

	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
