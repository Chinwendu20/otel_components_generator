package extensions

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/Chinwendu20/otel_components_generator/config"
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
)

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Funcs(template.FuncMap{
		"ObtainPackageName": func(module string) string {
			return strings.Split(module, "/")[2]

		},
	}).Parse(string(bytes)))
}
