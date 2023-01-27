package receivers

import (
	"text/template"
)

var templateSlice = []*template.Template{
	configTemplate,
	factoryTemplate,
	goModTemplate,
}
