package typing_test

import (
	"reflect"
	"testing"

	"example.com/ex00/typing"
)

// func TestGetWordAsset(t *testing.T) {
// 	type args struct {
// 		filePath string
// 	}
// 	tests := []struct {
// 		name      string
// 		args      args
// 		wantWords []string
// 		wantErr   bool
// 	}{
// 		// TODO: Add test cases.
// 		{"normal", args{"../testdata/test.txt"}, []string{"apple", "banana", "grape"}, false},
// 		{"normal", args{"../testdata/test2.file"}, []string{""}, false},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			gotWords, err := typing.GetWordAsset(tt.args.filePath)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetWordAsset() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotWords, tt.wantWords) {
// 				t.Errorf("GetWordAsset() = %v, want %v", gotWords, tt.wantWords)
// 			}
// 		})
// 	}
// }
