package extensions

import (
	_ "embed"
	"github.com/Chinwendu20/otel_components_generator/config"
	"strings"
	"text/template"
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
		"SplitString": func(signal string) []string {
			return strings.Split(signal, ",")
		},
	}).Parse(string(bytes)))
}
