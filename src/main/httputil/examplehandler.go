package httputil

import (
	"net/http"
	"gopractice/src/main/stringutil"
)

type Example struct {

}

func(ex *Example) ServeHTTP(re http.ResponseWriter, req *http.Request){
	re.Write(stringutil.Str2Byte(""))
}
