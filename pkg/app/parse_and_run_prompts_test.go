package app

import (
	"errors"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/prompt"
	"github.com/pasdam/go-scaffold/pkg/promptcli"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_parseAndRunPrompts(t *testing.T) {
	type mocks struct {
		prompts  []*prompt.Entry
		parseErr error
	}
	type args struct {
		promptsConfigPath string
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
		want  map[string]interface{}
	}{
		{
			name: "Should return data if no error occurs",
			mocks: mocks{
				prompts: []*prompt.Entry{{Name: "a"}, {Name: "b"}},
			},
			args: args{
				promptsConfigPath: "some-no-error-path",
			},
			want: map[string]interface{}{
				"ak": "av",
				"bk": "bv",
			},
		},
		{
			name: "Should handle parse error",
			mocks: mocks{
				parseErr: errors.New("some-parse-error"),
			},
			args: args{
				promptsConfigPath: "some-error-path",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			errHandler := func(v ...interface{}) {
				called = true
				assert.Equal(t, 2, len(v))
				assert.Equal(t, "Unable to parse prompts.yaml file. ", v[0])
				assert.Equal(t, tt.mocks.parseErr, v[1])
			}
			mockit.MockFunc(t, prompt.ParsePrompts).With(tt.args.promptsConfigPath).Return(tt.mocks.prompts, tt.mocks.parseErr)
			mockit.MockFunc(t, promptcli.RunPrompts).With(tt.mocks.prompts).Return(tt.want)

			got := parseAndRunPrompts(tt.args.promptsConfigPath, errHandler)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.mocks.parseErr != nil, called)
		})
	}
}
