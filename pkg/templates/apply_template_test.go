package templates

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func Test_applyTemplate_Fail_ShouldReturnErrorIfItFailsToExecuteTheTemplate(t *testing.T) {
	funcMap := template.FuncMap{}
	result, err := applyTemplate("This is a {{ .NotExistingProperty }}", "invalid_config", funcMap)

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func Test_applyTemplate_Fail_ShouldReturnErrorIfTemplateIsInvalid(t *testing.T) {
	funcMap := template.FuncMap{}
	data := struct{ CustomProperty string }{CustomProperty: "*test*"}
	result, err := applyTemplate("This is a {{ .CustomProperty } with invalid template", data, funcMap)

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func Test_applyTemplate_Success_ShouldCorrectlyGenerateOutputContentFromTemplate(t *testing.T) {
	funcMap := template.FuncMap{
		"Bold": func(value string) string {
			return fmt.Sprintf("*%s*", value)
		},
	}
	data := struct{ CustomProperty string }{CustomProperty: "test"}
	result, err := applyTemplate("This is a {{ Bold .CustomProperty }}", data, funcMap)

	assert.Nil(t, err)
	assert.Equal(t, "This is a *test*", result)
}
