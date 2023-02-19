//go:build tools
// +build tools

package tools

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/addlicense"
	_ "github.com/ory/go-acc"
	_ "github.com/pavius/impi/cmd/impi"
	_ "github.com/wadey/gocovmerge"
	_ "go.opentelemetry.io/build-tools/checkdoc"
	_ "go.opentelemetry.io/build-tools/chloggen"
	_ "go.opentelemetry.io/build-tools/multimod"
	_ "golang.org/x/tools/cmd/goimports"
)
