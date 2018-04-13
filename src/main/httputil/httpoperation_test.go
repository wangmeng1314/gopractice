package httputil

import "testing"

func TestHandleRequest(t *testing.T) {
	HandleRequest()
}

func TestSendGetRequest(t *testing.T) {
	res:=SendGetRequest("http://localhost:8080/info/test")
	println(res)
}
