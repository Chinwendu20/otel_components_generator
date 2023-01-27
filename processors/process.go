package processors

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"go.uber.org/zap"
	"text/template"
)

func GenerateProcessor(cfg config.ConfigStruct) []*template.Template {
	for _, signal := range cfg.SetSignals() {

		if signal == "metric" {
			templateSlice = append(templateSlice, metricTemplate)
		}
		if signal == "log" {
			templateSlice = append(templateSlice, logTemplate)
		}
		if signal == "trace" {
			templateSlice = append(templateSlice, traceTemplate)
		}
	}

	cfg.Logger.Info("Processor templates generated", zap.String("exporter", cfg.Module))

	return templateSlice
}
