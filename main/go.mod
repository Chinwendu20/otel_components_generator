module github.com/Chinwendu20/otel_components_generator

go 1.19

replace (
	github.com/Chinwendu20/otel_components_generator/config => ../config
	github.com/Chinwendu20/otel_components_generator/exporters => ../exporters
	github.com/Chinwendu20/otel_components_generator/extensions => ../extensions
	github.com/Chinwendu20/otel_components_generator/processors => ../processors
	github.com/Chinwendu20/otel_components_generator/receivers => ../receivers
)

require (
	github.com/Chinwendu20/otel_components_generator/config v0.0.0-20230125175729-930d7a4197e6
	github.com/Chinwendu20/otel_components_generator/exporters v0.0.0-00010101000000-000000000000
	github.com/Chinwendu20/otel_components_generator/extensions v0.0.0-00010101000000-000000000000
	github.com/Chinwendu20/otel_components_generator/processors v0.0.0-00010101000000-000000000000
	github.com/Chinwendu20/otel_components_generator/receivers v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	go.uber.org/multierr v1.9.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
