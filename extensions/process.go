package extensions

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"go.uber.org/zap"
	"text/template"
)

func GenerateExtension(cfg config.ConfigStruct) []*template.Template {

	cfg.Logger.Info("Extension templates generated", zap.String("extension", cfg.Module))

	return []*template.Template{
		configTemplate,
		configTestTemplate,
		factoryTemplate,
		factoryTestTemplate,
		goModTemplate,
	}
}
