package app

import (
	"errors"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_readOptions(t *testing.T) {
	type mocks struct {
		parseErr error
	}
	tests := []struct {
		name  string
		mocks mocks
		want  *config.Options
	}{
		{
			name:  "Should return parsed options",
			mocks: mocks{},
			want: &config.Options{
				OutputPath:       "some-out-path",
				TemplateRootPath: "some-template-path",
			},
		},
		{
			name: "Should handle error",
			mocks: mocks{
				parseErr: errors.New("some-parse-error"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockit.MockFunc(t, config.ParseCLIOptions).With().Return(tt.want, tt.mocks.parseErr)
			called := false
			errHandler := func(v ...interface{}) {
				called = true
				assert.Equal(t, 2, len(v))
				assert.Equal(t, "Command line options error. ", v[0])
				assert.Equal(t, tt.mocks.parseErr, v[1])
			}

			got := readOptions(errHandler)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.mocks.parseErr != nil, called)
		})
	}
}
