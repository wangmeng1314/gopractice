package httputil

import (
	"net/http"
	"io/ioutil"
	"gopractice/stringutil"
	"strings"
)

type Client struct {
	url    string ""
	method string ""
}

//发送GET请求，返回Body中的数据
func (*Client) SendGetRequest(url string) (string,http.Header){
	resp, err := http.Get(url)
	if nil != err {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode!=200 {
		return "",resp.Header
	}
	return stringutil.Byte2Str(body),resp.Header
}

func (*Client) SendGetRequestWithResp(url string) (string,http.Header){
	resp, err := http.Get(url)
	if nil != err {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	headers:=resp.Header

	return stringutil.Byte2Str(body),headers
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
func (*Client) SendRequest(url string, method string) string {
	client := http.DefaultClient

	request, err := http.NewRequest(method, url, strings.NewReader(""))
	handleErr(err)

	client.Do(request)

	return "success"
}

//捕捉错误
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
