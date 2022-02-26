package typing

import "testing"

func TestStartGame(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"normal", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartGame(); (err != nil) != tt.wantErr {
				t.Errorf("StartGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
