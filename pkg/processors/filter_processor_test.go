package processors

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestNewFilterProcessor(t *testing.T) {
	type args struct {
		filter        core.Filter
		nextProcessor core.Processor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				filter:        &mockFilter{true},
				nextProcessor: &mockProcessor{err: errors.New("some-err")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFilterProcessor(tt.args.filter, tt.args.nextProcessor).(*filterProcessor)

			assert.Equal(t, tt.args.filter, got.filter)
			assert.Equal(t, tt.args.nextProcessor, got.nextProcessor)
		})
	}
}

func Test_filterProcessor_ProcessFile(t *testing.T) {
	type args struct {
		path   string
		reader io.Reader
	}
	tests := []struct {
		name           string
		args           args
		shouldCallNext bool
		wantErr        error
	}{
		{
			name: "Should not call next processor if the path does not matche the filter",
			args: args{
				path:   "some-not-matching-path",
				reader: strings.NewReader("some-not-matching-reader"),
			},
			shouldCallNext: false,
			wantErr:        nil,
		},
		{
			name: "Should return error if the filter matches the path and the next processor raise it",
			args: args{
				path:   "some-next-processor-error-path",
				reader: strings.NewReader("some-next-processor-error-reader"),
			},
			shouldCallNext: true,
			wantErr:        errors.New("some-next-processor-error"),
		},
		{
			name: "Should not return error if the filter matches the path and the next processor does not raise it",
			args: args{
				path:   "some-next-processor-success-path",
				reader: strings.NewReader("some-next-processor-success-reader"),
			},
			shouldCallNext: true,
			wantErr:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			np := &mockProcessor{err: tt.wantErr}
			p := &filterProcessor{
				filter:        &mockFilter{tt.shouldCallNext},
				nextProcessor: np,
			}
			err := p.ProcessFile(tt.args.path, tt.args.reader)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.shouldCallNext, np.processed)
		})
	}
}
