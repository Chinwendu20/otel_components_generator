module github.com/Chinwendu20/otel_components_generator/extensions

go 1.19

replace github.com/Chinwendu20/otel_components_generator/config => ../config

require (
	github.com/Chinwendu20/otel_components_generator/config v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.24.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
