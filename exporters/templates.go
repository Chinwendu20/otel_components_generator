package exporters

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed templates/config.go.tmpl
	configBytes    []byte
	configTemplate = parseTemplate("components.go", configBytes)

	//go:embed templates/factory.go.tmpl
	factoryBytes    []byte
	factoryTemplate = parseTemplate("components_test.go", factoryBytes)

	//go:embed templates/go.mod.tmpl
	goModBytes    []byte
	goModTemplate = parseTemplate("main.go", goModBytes)

	//go:embed templates/log.go.tmpl
	logBytes    []byte
	logTemplate = parseTemplate("main_others.go", logBytes)

	//go:embed templates/metric.go.tmpl
	metricBytes    []byte
	metricTemplate = parseTemplate("main_windows.go", metricBytes)

	//go:embed templates/trace.go.tmpl
	traceBytes    []byte
	traceTemplate = parseTemplate("go.mod", traceBytes)
)

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Parse(string(bytes)))
}
