package httputil

import (
	"net/http"
	"io/ioutil"
	"gopractice/src/main/stringutil"
	"strings"
)

type Client struct {
	url    string ""
	method string ""
}

//发送GET请求，返回Body中的数据
func (*Client) SendGetRequest(url string) string {
	resp, err := http.Get(url)
	if nil != err {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return stringutil.Byte2Str(body)
}

//发送POST请求，返回Body中的数据
func (*Client) SendPostRequest(url string) string {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(""))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return stringutil.Byte2Str(body)
}

//发送HTTP请求
func (*Client) SendRequest(url string, method string) {
	client := http.DefaultClient

	request, err := http.NewRequest(method, url, strings.NewReader(""))
	handleErr(err)

	res, err := client.Do(request)
	handleErr(err)

	bodyByte, err := ioutil.ReadAll(res.Body)
	handleErr(err)

	println(stringutil.Byte2Str(bodyByte))
}

//捕捉错误
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
