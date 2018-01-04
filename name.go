package build

import (
	"strings"
	"text/template"
)

const defaultPostfix = "_{{.OS}}-{{.Arch}}{{.Extension}}"

// DefaultNameTemplate returns a template for an executable name
// consisting of the baseName followed by the OS, architecture and
// optional extension.
func DefaultNameTemplate(baseName string) *template.Template {
	tmpl := template.New("executableName")
	return template.Must(tmpl.Parse(escapeName(baseName) + defaultPostfix))
}

func escapeName(baseName string) string {
	baseName = strings.Replace(baseName, "{{", "{{\"{{\"}}", -1)
	baseName = strings.Replace(baseName, "}}", "{{\"}}\"}}", -1)
	return baseName
}
