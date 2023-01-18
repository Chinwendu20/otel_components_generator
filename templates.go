package Otel_Component_Generator

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed templates/components.go.tmpl
	componentsBytes    []byte
	componentsTemplate = parseTemplate("components.go", componentsBytes)

	//go:embed templates/components_test.go.tmpl
	componentsTestBytes    []byte
	componentsTestTemplate = parseTemplate("components_test.go", componentsTestBytes)

	//go:embed templates/main.go.tmpl
	mainBytes    []byte
	mainTemplate = parseTemplate("main.go", mainBytes)

	//go:embed templates/main_others.go.tmpl
	mainOthersBytes    []byte
	mainOthersTemplate = parseTemplate("main_others.go", mainOthersBytes)

	//go:embed templates/main_windows.go.tmpl
	mainWindowsBytes    []byte
	mainWindowsTemplate = parseTemplate("main_windows.go", mainWindowsBytes)

	//go:embed templates/go.mod.tmpl
	goModBytes    []byte
	goModTemplate = parseTemplate("go.mod", goModBytes)
)

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Parse(string(bytes)))
}
