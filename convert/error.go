package convert

import (
	"errors"
	"fmt"
)

type ErrCode string

var (
	InValidSrcDirPath ErrCode = "invalid src directory path"
	InValidDstDirPath ErrCode = "invalid dst directory path"
	InValidExt        ErrCode = "invalid file extension"
	FileAccessFail    ErrCode = "cannot access file"
	FileOpenFail      ErrCode = "cannot open file"
	ImgCreateFail     ErrCode = "cannot create image from file"
	InValidOutputPath ErrCode = "cannot get output filepath for"
	FileOutputFail    ErrCode = "cannot create output file"
	FileEncodeFail    ErrCode = "cannot encode img file"
)

var (
	ErrExt = errors.New("invalid file extension")
)

type ConvError struct {
	Err      error
	Code     ErrCode
	FilePath string
}

func (e *ConvError) Error() string {
	return fmt.Sprintln(e.Code, e.FilePath)
}

func (e *ConvError) Unwrap() error { return e.Err }
