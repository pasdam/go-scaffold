package templates

import (
	"bytes"
	"text/template"
)

func applyTemplate(templateContent string, config interface{}, funcMap template.FuncMap) (string, error) {
	template, err := template.New("").Funcs(funcMap).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = template.Execute(&result, config)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
