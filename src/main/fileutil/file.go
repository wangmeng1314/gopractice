package fileutil

import (
	"io/ioutil"
	"gopractice/src/main/stringutil"
	"os"
)

type Reader interface {
	readFile() string
}

type Writer interface {
	writeFile(path string,data string)
}

type FileProcessor struct {
	path string
}
//读取文件内容
func (fprocessor *FileProcessor) readFile() string {
	resByte, err := ioutil.ReadFile(fprocessor.path)

	if err != nil {
		panic(err)
	}

	res := stringutil.Byte2Str(resByte)

	return res
}
//写入文件
func (fprocessor *FileProcessor) writeFile(path string,data string) {
	ioutil.WriteFile(path,stringutil.Str2Byte(data),os.ModeDevice)
}
