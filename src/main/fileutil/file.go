package fileutil

import (
	"io/ioutil"
	"gopractice/src/main/stringutil"
	"os"
	"bufio"
	"io"
	"strings"
	"net/url"
	"fmt"
	"gopractice/src/main/httputil"
	"container/list"
	"log"
)

type Reader interface {
	readFile() string
}

type Writer interface {
	writeFile(path string, data string)
}

type FileProcessor struct {
	path string
}

//读取文件内容
func (fprocessor *FileProcessor) ReadFile() string {
	resByte, err := ioutil.ReadFile(fprocessor.path)

	if err != nil {
		panic(err)
	}

	res := stringutil.Byte2Str(resByte)

	return res
}

//写入文件
func (fprocessor *FileProcessor) writeFile(path string, data string) {
	ioutil.WriteFile(path, stringutil.Str2Byte(data), os.ModeDevice)
}

//按行读取文件
func (fprocessor *FileProcessor) ReadFilePerLine() string {
	file, err := os.Open(fprocessor.path)
	client := httputil.Client{}
	buffer := bufio.NewReader(file)
	for {
		path := "https://campaigncdn.m.taobao.com/shop/campaign-cdnpage.htm?"
		res, _, c := buffer.ReadLine()
		if c == io.EOF {
			break
		}
		a := stringutil.Byte2Str(res)
		res1 := strings.Split(a, ",")
		for e := range res1 {
			if e == 1 {
				path = path + "userId=" + res1[e]
			}
			if e == 2 {
				//解析URL
				url, _ := url.Parse(res1[e])
				path = path + "&pathInfo=" + url.Path
			}
		}
		fmt.Println(path)
		httpRes, _ := client.SendGetRequest(path)
		contains := strings.Contains(httpRes, "globalData")
		if !contains {
			fmt.Println("fail url", path)
		}
	}

	resByte, err := ioutil.ReadFile(fprocessor.path)

	if err != nil {
		panic(err)
	}

	res := stringutil.Byte2Str(resByte)

	return res
}

func (fprocessor *FileProcessor) ReadFilePerLineAsList() *list.List {
	file, err := os.Open(fprocessor.path)

	if err != nil {
		panic(err)
	}

	client := httputil.Client{}
	urlList := list.New()
	buffer := bufio.NewReader(file)
	sucessCount := 0
	newCount := 0
	for {
		res, _, c := buffer.ReadLine()
		if c == io.EOF {
			break
		}
		path := stringutil.Byte2Str(res)

		urlList.PushBack(path)

		bodyRes, header := client.SendGetRequest(path)

		if strings.Contains(bodyRes, "哎呦喂，被挤爆了") {
			log.Print("fail url" + path)
		}

		age := header.Get("age")
		log.Print("success----" + age)
		if age == "" {
			newCount += 1
		}
		sucessCount += 1
	}

	log.Print("共预热", urlList.Len(), "个URL")
	log.Print("成功", sucessCount, "个URL")
	log.Print("失败", urlList.Len()-sucessCount, "个URL")
	log.Print("新预热", newCount, "个URL")

	return urlList
}

func (fprocessor *FileProcessor) ReadFilePerLineUA() {
	file, err := os.Open(fprocessor.path)
	buffer := bufio.NewReader(file)
	i := 1
	for {
		i++
		res, _, c := buffer.ReadLine()
		if c == io.EOF {
			break
		}

		uaRes := stringutil.Byte2Str(res)
		uaSpiltRes := strings.Split(uaRes, "****")
		for e := range uaSpiltRes {
			if e == 1 {
				ua := uaSpiltRes[e]
				fmt.Println(ua)
			}
		}

		if strings.Contains(stringutil.Byte2Str(res), "MISS") {
			fmt.Println(stringutil.Byte2Str(res))
		}
	}

	if err != nil {
		panic(err)
	}
}
