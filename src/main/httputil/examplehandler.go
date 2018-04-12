package httputil

import (
	"net/http"
)

type Example struct {

}

func(ex *Example) ServeHTTP(re http.ResponseWriter, req *http.Request){

}
