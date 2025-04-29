package cli

import (
	"testing"
)

func TestShowOptionDescription(t *testing.T) {
	for i := range 7 {
		desc, err := showOptionDescription(i)
		if err != nil {
			t.Errorf("unsupported option %d", i)
		} else {
			t.Logf("test option %d passed: %s", i, desc)
		}
	}
}

func TestGetNumber(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
