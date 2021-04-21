package templates

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"github.com/pasdam/mockit/matchers/argument"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_ProcessTemplate_fail_shouldPropagateErrorIfReaderThrowsIt(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()
	funcMap := template.FuncMap{}
	expectedErr := errors.New("some-read-error")
	mockit.MockMethod(t, file, file.Read).With(argument.Any).Return(0, expectedErr)

	reader, err := ProcessTemplate(file, "invalid-data", funcMap)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, reader)
}

func Test_ProcessTemplate_fail_shouldReturnErrorIfApplyingTheTemplateFailed(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()
	funcMap := template.FuncMap{}

	reader, err := ProcessTemplate(file, "invalid-data", funcMap)

	assert.NotNil(t, err)
	assert.Nil(t, reader)
}

func Test_ProcessTemplate_success_shouldCreateAReaderForTheGeneratedContent(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()
	funcMap := template.FuncMap{
		"Bold": func(value string) string {
			return fmt.Sprintf("*%s*", value)
		},
	}

	reader, err := ProcessTemplate(file, struct{ Text string }{Text: "test"}, funcMap)

	assert.Nil(t, err)
	readContent, err := ioutil.ReadAll(reader)
	assert.Nil(t, err)
	assert.Equal(t, "This is a *test*\n", string(readContent))
}
