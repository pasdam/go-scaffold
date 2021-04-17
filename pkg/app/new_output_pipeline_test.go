package app

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_newOutputPipeline_ShouldProcessFileAsExpected(t *testing.T) {
	data := make(map[string]interface{})
	data["text"] = "test!"
	outDir := filestest.TempDir(t)
	type args struct {
		path    string
		content string
	}
	type expect struct {
		shouldExist bool
		path        string
	}
	tests := []struct {
		name   string
		args   args
		expect expect
	}{
		{
			name: "Should process files with .go-scaffold prefix",
			args: args{
				path:    ".go-scaffold-file",
				content: "",
			},
			expect: expect{
				shouldExist: true,
				path:        ".go-scaffold-file",
			},
		},
		{
			name: "Should process .git folder",
			args: args{
				path:    filepath.Join(".git", "some-git-file"),
				content: "",
			},
			expect: expect{
				shouldExist: true,
				path:        filepath.Join(".git", "some-git-file"),
			},
		},
		{
			name: "Should process files with .git prefix",
			args: args{
				path:    ".gitignore",
				content: "",
			},
			expect: expect{
				shouldExist: true,
				path:        ".gitignore",
			},
		},
		{
			name: "Should process go-scaffold files",
			args: args{
				path:    filepath.Join(".go-scaffold", "prompts.yaml"),
				content: "",
			},
			expect: expect{
				shouldExist: true,
				path:        filepath.Join(".go-scaffold", "prompts.yaml"),
			},
		},
		{
			name: "Should process regular files",
			args: args{
				path:    "some-regular-file.txt",
				content: "some-regular-file-content",
			},
			expect: expect{
				shouldExist: true,
				path:        "some-regular-file.txt",
			},
		},
		{
			name: "Should process template files",
			args: args{
				path:    "some-template-file.txt.tpl",
				content: "some-template-file-content",
			},
			expect: expect{
				shouldExist: true,
				path:        "some-template-file.txt",
			},
		},
		{
			name: "Should process go-scaffold config files",
			args: args{
				path:    "some-template-file.txt.tpl",
				content: "some-template-file-content",
			},
			expect: expect{
				shouldExist: true,
				path:        "some-template-file.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newOutputPipeline(data, outDir, &scaffold.TemplateHelper{})
			assert.NotNil(t, got)
			assert.Nil(t, err)

			reader := strings.NewReader(tt.args.content)

			err = got.ProcessFile(tt.args.path, reader)
			assert.Nil(t, err)

			outPath := filepath.Join(outDir, tt.expect.path)
			if tt.expect.shouldExist {
				filestest.FileExistsWithContent(t, outPath, tt.args.content)
			} else {
				filestest.PathDoesNotExist(t, outPath)
			}
		})
	}
}
