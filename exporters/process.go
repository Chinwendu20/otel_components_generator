package exporters

import (
	"text/template"

	"go.uber.org/zap"

	"github.com/Chinwendu20/otel_components_generator/config"
)

// Generates templates for processors
func GenerateExporter(cfg config.Struct) []*template.Template {
	templateSlice := []*template.Template{
		configTemplate,
		configTestTemplate,
		factoryTemplate,
		factoryTestTemplate,
		goModTemplate,
	}

	for _, signal := range cfg.SetSignals() {

		if signal == "metric" {
			templateSlice = append(templateSlice, metricTemplate, metricTestTemplate)
		}
		if signal == "log" {
			templateSlice = append(templateSlice, logTemplate, logTestTemplate)
		}
		if signal == "trace" {
			templateSlice = append(templateSlice, traceTemplate, traceTestTemplate)
		}
	}

	cfg.Logger.Info("Exporter templates generated", zap.String("exporter", cfg.Module))

	return templateSlice
}
