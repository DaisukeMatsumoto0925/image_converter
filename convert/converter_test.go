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
			name: "png to jpg",
			args: args{srcDir: "./images", dstDir: "./result", bExt: "png", aExt: "jpg"},
			want: &Converter{
				srcDirPath: "/Users/matsumotodaisuke/projects/public/image_converter/convert/images",
				dstDirPath: "/Users/matsumotodaisuke/projects/public/image_converter/convert/result",
				bExt:       "png",
				aExt:       "jpg",
			},
			wantErr: false,
		},
		{
			name: "jpg to png",
			args: args{srcDir: "./images", dstDir: "./result", bExt: "jpg", aExt: "png"},
			want: &Converter{
				srcDirPath: "/Users/matsumotodaisuke/projects/public/image_converter/convert/images",
				dstDirPath: "/Users/matsumotodaisuke/projects/public/image_converter/convert/result",
				bExt:       "jpg",
				aExt:       "png",
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

func TestConverter_convert(t *testing.T) {
	type fields struct {
		srcDirPath string
		dstDirPath string
		bExt       string
		aExt       string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				srcDirPath: tt.fields.srcDirPath,
				dstDirPath: tt.fields.dstDirPath,
				bExt:       tt.fields.bExt,
				aExt:       tt.fields.aExt,
			}
			if err := c.convert(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Converter.convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConverter_getOutputFileName(t *testing.T) {
	type fields struct {
		srcDirPath string
		dstDirPath string
		bExt       string
		aExt       string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				srcDirPath: tt.fields.srcDirPath,
				dstDirPath: tt.fields.dstDirPath,
				bExt:       tt.fields.bExt,
				aExt:       tt.fields.aExt,
			}
			got, err := c.getOutputFileName(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Converter.getOutputFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Converter.getOutputFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_Do(t *testing.T) {
	type fields struct {
		srcDirPath string
		dstDirPath string
		bExt       string
		aExt       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{
				srcDirPath: tt.fields.srcDirPath,
				dstDirPath: tt.fields.dstDirPath,
				bExt:       tt.fields.bExt,
				aExt:       tt.fields.aExt,
			}
			if err := c.Do(); (err != nil) != tt.wantErr {
				t.Errorf("Converter.Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
