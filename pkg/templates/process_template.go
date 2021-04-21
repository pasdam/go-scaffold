package templates

import (
	"io"
	"log"
	"strings"
	"text/template"
)

// ProcessTemplate processes the template using the specified data
func ProcessTemplate(reader io.Reader, data interface{}, funcMap template.FuncMap) (io.Reader, error) {
	byteContent, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	content, err := applyTemplate(string(byteContent), data, funcMap)
	if err != nil {
		log.Println("Error while generating output file from template")
		return nil, err
	}

	return strings.NewReader(content), nil
}
