package exporters

import (
	"github.com/Chinwendu20/otel_components_generator/internal"
	"go.uber.org/zap"
	"text/template"
)

func GenerateExporter(cfg internal.ConfigStruct) []*template.Template {
	for _, signal := range internal.Config.Signals {

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

	internal.Config.Logger.Info("Exporter templates generated", zap.String("exporter", internal.Config.Module))

	return templateSlice
}
