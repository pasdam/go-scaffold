package processors

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func TestNewWriteProcessor(t *testing.T) {
	tests := []struct {
		name string
		want core.Processor
	}{
		{
			name: "Should create instance",
			want: &writeProcessor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWriteProcessor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriteProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeProcessor_ProcessFile(t *testing.T) {
	type mocks struct {
		writeErr error
	}
	type args struct {
		filePath string
		reader   io.Reader
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
	}{
		{
			name: "Should return error if ioutilx.ReaderToFile raises it",
			mocks: mocks{
				writeErr: errors.New("some-write-error"),
			},
			args: args{
				filePath: "some-write-error-path",
				reader:   strings.NewReader("some-write-error-reader"),
			},
		},
		{
			name: "Should not return error if ioutilx.ReaderToFile succeed",
			args: args{
				filePath: "some-success-path",
				reader:   strings.NewReader("some-success-reader"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := tt.mocks.writeErr
			mockit.MockFunc(t, ioutilx.ReaderToFile).With(tt.args.reader, tt.args.filePath).Return(wantErr)
			p := &writeProcessor{}

			err := p.ProcessFile(tt.args.filePath, tt.args.reader)

			assert.Equal(t, wantErr, err)
		})
	}
}
