package exporters

import (
	_ "embed"
	"strings"
	"text/template"
)

var (
	//go:embed templates/config.go.tmpl
	configBytes    []byte
	configTemplate = parseTemplate("config.go", configBytes)

	//go:embed templates/factory.go.tmpl
	factoryBytes    []byte
	factoryTemplate = parseTemplate("factory.go", factoryBytes)

	//go:embed templates/go.mod.tmpl
	goModBytes    []byte
	goModTemplate = parseTemplate("go.mod", goModBytes)

	//go:embed templates/log.go.tmpl
	logBytes    []byte
	logTemplate = parseTemplate("log.go", logBytes)

	//go:embed templates/metric.go.tmpl
	metricBytes    []byte
	metricTemplate = parseTemplate("metric.go", metricBytes)

	//go:embed templates/trace.go.tmpl
	traceBytes    []byte
	traceTemplate = parseTemplate("trace.go", traceBytes)
)

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Funcs(template.FuncMap{
		"SplitString": func(signal string) []string {
			return strings.Split(signal, ",")
		},
	}).Parse(string(bytes)))
}
