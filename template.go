package main

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	kvstoreProto "github.com/ehsundar/kvstore/protobuf/kvstore"
)

//go:embed kvstore.tmpl
var rawTemplate string

type kvstoreTemplateContext struct {
	PackageName string

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
}

type valueSpecs struct {
	Opts *kvstoreProto.KVStoreValueOptions

	MessageName string
}

func Render(templateContext kvstoreTemplateContext) (string, error) {
	tmpl := template.Must(
		template.New("base").Funcs(sprig.FuncMap()).Parse(rawTemplate),
	)

	builder := &bytes.Buffer{}
	err := tmpl.Execute(builder, templateContext)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
