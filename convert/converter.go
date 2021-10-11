package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const (
	PNG  = "png"
	JPG  = "jpg"
	JPEG = "jpeg"
)

var flagToExtNames map[string][]string = map[string][]string{
	PNG:  {".png"},
	JPG:  {".jpg", ".jpeg"},
	JPEG: {".jpeg", ".jpg"},
}

func contains(slice []string, elem string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}
	return false
}

type Converter struct {
	srcDirPath string
	dstDirPath string
	bExt       string
	aExt       string
}

func NewConverter(srcDir, dstDir, bExt, aExt string) (*Converter, error) {
	srcDirAbs, err := absPath(srcDir)
	if err != nil {
		return nil, &ConvError{Err: err, Code: InValidSrcDirPath, FilePath: srcDir}
	}

	dstDirAbs, err := absPath(dstDir)
	if err != nil {
		return nil, &ConvError{Err: err, Code: InValidSrcDirPath, FilePath: dstDir}
	}

	if _, ok := flagToExtNames[bExt]; !ok {
		return nil, &ConvError{Err: ErrExt, Code: InValidExt, FilePath: bExt}
	}

	if _, ok := flagToExtNames[aExt]; !ok {
		return nil, &ConvError{Err: ErrExt, Code: InValidExt, FilePath: aExt}
	}

	return &Converter{
		srcDirPath: srcDirAbs,
		dstDirPath: dstDirAbs,
		bExt:       bExt,
		aExt:       aExt,
	}, err
}

func (c *Converter) Do() error {
	err := filepath.Walk(c.srcDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return &ConvError{Err: err, Code: FileAccessFail, FilePath: path}
			}

			if info.IsDir() {
				return nil
			}

			if contains(flagToExtNames[c.bExt], filepath.Ext(path)) {
				err := c.convert(path)
				fmt.Println((path))
				if err != nil {
					return err
				}
			}
			return nil
		})

	if err != nil {
		return err
	}
	return nil
}

func (c *Converter) convert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return &ConvError{Err: err, Code: FileOpenFail, FilePath: path}
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return &ConvError{Err: err, Code: ImgCreateFail, FilePath: path}
	}

	newFileName, err := c.getOutputFileName(path)
	if err != nil {
		return &ConvError{Err: err, Code: InValidOutputPath, FilePath: path}
	}
	newFileDirName := filepath.Dir(newFileName)
	if err := os.MkdirAll(newFileDirName, 0777); err != nil {
		return &ConvError{Err: err, Code: InValidDstDirPath, FilePath: c.dstDirPath}
	}

	newfile, err := os.Create(newFileName)
	if err != nil {
		return &ConvError{Err: err, Code: FileOutputFail, FilePath: newFileName}
	}
	defer func() {
		err := newfile.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	switch c.aExt {
	case PNG:
		err = png.Encode(newfile, img)
		if err != nil {
			return &ConvError{Err: err, Code: FileEncodeFail, FilePath: newFileName}
		}
	case JPG, JPEG:
		err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: 75})
		if err != nil {
			return &ConvError{Err: err, Code: FileEncodeFail, FilePath: newFileName}
		}
	}
	return nil
}

func (c *Converter) getOutputFileName(path string) (string, error) {
	rel, err := filepath.Rel(c.srcDirPath, path)
	if err != nil {
		return "", err
	}
	fNameWithoutExt := removeFileExt(filepath.Join(c.dstDirPath, rel))

	newExt := flagToExtNames[c.aExt][0]

	return fNameWithoutExt + newExt, nil
}
