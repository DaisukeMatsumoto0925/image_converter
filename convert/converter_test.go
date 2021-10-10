package convert

import (
	"reflect"
	"testing"
)

func TestNewConverter(t *testing.T) {
	type args struct {
		srcDir string
		dstDir string
		bExt   string
		aExt   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Converter
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{srcDir: "./images", dstDir: "./result", bExt: "png", aExt: "jpg"},
			want: &Converter{
				srcDirPath: "/Users/matsumotodaisuke/projects/public/image_converter/convert/images",
				dstDirPath: "/Users/matsumotodaisuke/projects/public/image_converter/convert/result",
				bExt:       "png",
				aExt:       "jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConverter(tt.args.srcDir, tt.args.dstDir, tt.args.bExt, tt.args.aExt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
