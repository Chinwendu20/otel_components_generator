package extensions

import (
	"text/template"

	"go.uber.org/zap"

	"github.com/tejas-contentstack/otel_components_generator/config"
)

// Generates templates for processors
func GenerateExtension(cfg config.Struct) []*template.Template {

	cfg.Logger.Info("Extension templates generated", zap.String("extension", cfg.Module))

	return []*template.Template{
		configTemplate,
		configTestTemplate,
		factoryTemplate,
		factoryTestTemplate,
		goModTemplate,
	}
}
