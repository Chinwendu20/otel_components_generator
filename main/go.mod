module github.com/tejas-contentstack/otel_components_generator

go 1.25

replace (
	github.com/tejas-contentstack/otel_components_generator/config => ../config
	github.com/tejas-contentstack/otel_components_generator/exporters => ../exporters
	github.com/tejas-contentstack/otel_components_generator/extensions => ../extensions
	github.com/tejas-contentstack/otel_components_generator/processors => ../processors
	github.com/tejas-contentstack/otel_components_generator/receivers => ../receivers
)

require (
	github.com/tejas-contentstack/otel_components_generator/config v0.0.0-20230304090600-f77865683ced
	github.com/tejas-contentstack/otel_components_generator/exporters v0.0.0-20230304090600-f77865683ced
	github.com/tejas-contentstack/otel_components_generator/extensions v0.0.0-20230304090600-f77865683ced
	github.com/tejas-contentstack/otel_components_generator/processors v0.0.0-20230304090600-f77865683ced
	github.com/tejas-contentstack/otel_components_generator/receivers v0.0.0-20230304090600-f77865683ced
	github.com/spf13/cobra v1.10.2
	github.com/stretchr/testify v1.8.1
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
