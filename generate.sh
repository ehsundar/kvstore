#!/bin/bash

find . -type f -name "*.go"  -exec grep -qE "// Code generated by (mockery|sqlc|protoc-gen-go-kvstore).*\. DO NOT EDIT\." {} \; -delete

protoc \
  --go_out=. --go_opt=paths=source_relative \
  protobuf/kvstore/options.proto

go build -o protoc-gen-go-kvstore "$(pwd)/cmd/protoc-gen-go-kvstore"

protoc \
  --plugin=protoc-gen-go-kvstore=./protoc-gen-go-kvstore \
  --go_out=. --go_opt=paths=source_relative \
  --go-kvstore_out=. --go-kvstore_opt=paths=source_relative \
  -I. \
  -I./protobuf \
  examples/example.proto
