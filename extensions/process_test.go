package extensions

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
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
