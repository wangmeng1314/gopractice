package httputil

import (
	"net/http"
	"gopractice/stringutil"
)

type Example struct {

}

func(ex *Example) ServeHTTP(re http.ResponseWriter, req *http.Request){
	re.Write(stringutil.Str2Byte(""))
}
