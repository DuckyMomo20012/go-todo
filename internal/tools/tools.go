//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen" // OpenAPI 3.0 code generator
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // Linter
	_ "golang.org/x/tools/cmd/goimports"                    // goimports
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"        // Protobuf generator
	_ "honnef.co/go/tools/cmd/staticcheck"                  // Staticcheck
	_ "mvdan.cc/gofumpt"                                    // gofumpt
)
