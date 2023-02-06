package exporters

import (
	_ "embed"
	"github.com/Chinwendu20/otel_components_generator/config"
	"strings"
	"text/template"
)

var (
	//go:embed templates/config.go.tmpl
	configBytes    []byte
	configTemplate = parseTemplate(config.ConfigFileName, configBytes)

	//go:embed templates/config_test.go.tmpl
	configTestBytes    []byte
	configTestTemplate = parseTemplate(config.ConfigTestFileName, configTestBytes)

	//go:embed templates/factory.go.tmpl
	factoryBytes    []byte
	factoryTemplate = parseTemplate(config.FactoryFileName, factoryBytes)

	//go:embed templates/factory_test.go.tmpl
	factoryTestBytes    []byte
	factoryTestTemplate = parseTemplate(config.FactoryTestFileName, factoryTestBytes)

	//go:embed templates/go.mod.tmpl
	goModBytes    []byte
	goModTemplate = parseTemplate(config.GoModFileName, goModBytes)

	//go:embed templates/log.go.tmpl
	logBytes    []byte
	logTemplate = parseTemplate(config.LogFileName, logBytes)

	//go:embed templates/log_test.go.tmpl
	logTestBytes    []byte
	logTestTemplate = parseTemplate(config.LogTestFileName, logTestBytes)

	//go:embed templates/metric_test.go.tmpl
	metricTestBytes    []byte
	metricTestTemplate = parseTemplate(config.MetricTestFileName, metricTestBytes)

	//go:embed templates/metric.go.tmpl
	metricBytes    []byte
	metricTemplate = parseTemplate(config.MetricFileName, metricBytes)

	//go:embed templates/trace.go.tmpl
	traceBytes    []byte
	traceTemplate = parseTemplate(config.TraceFileName, traceBytes)

	//go:embed templates/trace_test.go.tmpl
	traceTestBytes    []byte
	traceTestTemplate = parseTemplate(config.TraceTestFileName, traceTestBytes)
)

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Funcs(template.FuncMap{
		"SplitString": func(signal string) []string {
			return strings.Split(signal, ",")
		},
		"ObtainPackageName": func(module string) string {
			return strings.Split(module, "/")[2]

		},
	}).Parse(string(bytes)))
}
