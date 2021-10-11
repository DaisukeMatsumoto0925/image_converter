package convert

import (
	"path/filepath"
	"reflect"
	"testing"
)

func getAbs(t *testing.T, path string) string {
	t.Helper()
	abs, err := filepath.Abs(path)
	if err != nil {
		t.Errorf("could not get abs path %v", path)
	}
	return abs

}

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
				srcDirPath: getAbs(t, "./images"),
				dstDirPath: getAbs(t, "result"),
				bExt:       "png",
				aExt:       "jpg",
			},
			wantErr: false,
		},
		{
			name: "jpg to png",
			args: args{srcDir: "./images", dstDir: "./result", bExt: "jpg", aExt: "png"},
			want: &Converter{
				srcDirPath: getAbs(t, "./images"),
				dstDirPath: getAbs(t, "./result"),
				bExt:       "jpg",
				aExt:       "png",
			},
			wantErr: false,
		},
		{
			name: "abs to rel path",
			args: args{
				srcDir: getAbs(t, "./images"),
				dstDir: "./result",
				bExt:   "jpg",
				aExt:   "png",
			},
			want: &Converter{
				srcDirPath: getAbs(t, "./images"),
				dstDirPath: getAbs(t, "./result"),
				bExt:       "jpg",
				aExt:       "png",
			},
			wantErr: false,
		},
		{
			name:    "from gif",
			args:    args{srcDir: "./images", dstDir: "./result", bExt: "gif", aExt: "png"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "to gif",
			args:    args{srcDir: "./images", dstDir: "./result", bExt: "png", aExt: "gif"},
			want:    nil,
			wantErr: true,
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
		{
			name: "png to jpg",
			fields: fields{
				srcDirPath: getAbs(t, "../testdata/images"),
				dstDirPath: getAbs(t, "../testdata/result"),
				bExt:       "png",
				aExt:       "jpg",
			},
			args:    args{path: getAbs(t, "../testdata/images/1.png")},
			wantErr: false,
		},
		{
			name: "jpg to png",
			fields: fields{
				srcDirPath: getAbs(t, "../testdata/images"),
				dstDirPath: getAbs(t, "../testdata/result"),
				bExt:       "jpg",
				aExt:       "png",
			},
			args:    args{path: getAbs(t, "../testdata/images/1.png")},
			wantErr: false,
		},
		{
			name: "fail path",
			fields: fields{
				srcDirPath: getAbs(t, "../testdata/images"),
				dstDirPath: getAbs(t, "../testdata/result"),
				bExt:       "jpg",
				aExt:       "png",
			},
			args:    args{path: getAbs(t, "./testdata/images/1.png")},
			wantErr: true,
		},
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
