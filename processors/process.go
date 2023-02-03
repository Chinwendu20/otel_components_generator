package processors

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"go.uber.org/zap"
	"text/template"
)

func GenerateProcessor(cfg config.ConfigStruct) []*template.Template {
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
	cfg.Logger.Info("Processor templates generated", zap.String("processor", cfg.Module))

	return templateSlice
}
