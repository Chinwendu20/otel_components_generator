package extensions

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"

	"github.com/tejas-contentstack/otel_components_generator/config"
)

func TestGenerateExtension(t *testing.T) {
	cfg := config.NewConfig()
	assert.Equal(t, len(GenerateExtension(cfg)), len([]*template.Template{
		configTemplate,
		configTestTemplate,
		factoryTemplate,
		factoryTestTemplate,
		goModTemplate,
	}))
}
