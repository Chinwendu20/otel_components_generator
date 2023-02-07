package receivers

import (
	"text/template"

	"go.uber.org/zap"

	"github.com/Chinwendu20/otel_components_generator/config"
)

func GenerateReceiver(cfg config.ConfigStruct) []*template.Template {
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

	cfg.Logger.Info("Receiver templates generated", zap.String("exporter", cfg.Module))

	return templateSlice
}
