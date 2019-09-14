package testutils

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FileDoesNotExist is a test helper function used to assert that the specified file does not exist
func FileDoesNotExist(t *testing.T, filePath string) {
	_, err := os.Stat(filePath)
	assert.NotNil(t, err)
	assert.True(t, strings.HasSuffix(err.Error(), ": no such file or directory"))
}
