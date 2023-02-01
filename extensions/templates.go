package extensions

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed templates/config.go.tmpl
	configBytes    []byte
	configTemplate = parseTemplate("config.go", configBytes)

	//go:embed templates/factory.go.tmpl
	factoryBytes    []byte
	factoryTemplate = parseTemplate("factory.go", factoryBytes)

	//go:embed templates/go.mod.tmpl
	goModBytes    []byte
	goModTemplate = parseTemplate("go.mod", goModBytes)
)

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Parse(string(bytes)))
}
