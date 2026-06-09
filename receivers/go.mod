module github.com/tejas-contentstack/otel_components_generator/receivers

go 1.25

replace github.com/tejas-contentstack/otel_components_generator/config => ../config

require (
	github.com/tejas-contentstack/otel_components_generator/config v0.0.0-20230304090600-f77865683ced
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
