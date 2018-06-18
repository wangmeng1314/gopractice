package fileutil

import (
	"testing"
	"time"
	"fmt"
)

func TestFileProcessor_ReadFilePerLine(t *testing.T) {
	f := FileProcessor{"/Users/fulei.yang/go/src/gopractice/src/main/fileutil/file.txt"}
	f.ReadFilePerLine()
}

func TestFileProcessor_ReadFilePerLineAsList(t *testing.T) {
	f := FileProcessor{"/Users/fulei.yang/go/src/gopractice/src/main/fileutil/url.txt"}
	for ; ; {
		fmt.Println(time.Now())
		fmt.Println("waiting for preheat")
		//time.Sleep(time.Minute * 20)
		f.ReadFilePerLineAsList()

	}
}

func TestFileProcessor_ReadFilePerLineUA(t *testing.T) {
	f := FileProcessor{"/Users/fulei.yang/go/src/gopractice/src/main/fileutil/ua.txt"}
	f.ReadFilePerLineUA()
}
