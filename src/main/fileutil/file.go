package fileutil

import (
	"io/ioutil"
	"gopractice/src/main/stringutil"
)

type Reader interface {
	readFile() string
}

type FileProcessor struct {
	path string
}

func (fprocessor *FileProcessor) readFile() string {
	resByte, err := ioutil.ReadFile(fprocessor.path)

	if err != nil {
		panic(err)
	}

	res := stringutil.Byte2Str(resByte)
	return res

}
