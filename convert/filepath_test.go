package convert

import (
	"testing"
)

func Test_removeFileExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeFileExt(tt.args.path); got != tt.want {
				t.Errorf("removeFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := absPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("absPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("absPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
