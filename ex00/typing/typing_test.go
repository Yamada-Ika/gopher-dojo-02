package typing_test

import (
	"bytes"
	"testing"

	"example.com/ex00/typing"
)

func TestExportPrintProblem(t *testing.T) {
	type args struct {
		w    *bytes.Buffer
		word string
	}
	tests := []struct {
		name string
		args args
		exp  string
	}{
		// TODO: Add test cases.
		{"normal", args{&bytes.Buffer{}, "hoge"}, "hoge\n->"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			typing.ExportPrintProblem(tt.args.w, tt.args.word)
			if tt.args.w.String() != tt.exp {
				t.Errorf("ExportPrintProblem() got = %v, expected %v", tt.args.w.String(), tt.exp)
			}
		})
	}
}
