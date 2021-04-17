package filters

import "testing"

func Test_noOpFilter_Accept(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should accept any value",
			args: args{
				value: "some-value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewNoOpFilter()
			if got := f.Accept(tt.args.value); got != true {
				t.Errorf("noOpFilter.Accept() = %v, want %v", got, true)
			}
		})
	}
}
