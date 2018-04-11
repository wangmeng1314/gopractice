package httputil

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleRequest() {
	rounter := mux.Router{}
	rounter.HandleFunc("/info/{action}", HandlePost).Methods("GET")
}

func HandlePost(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var str string = "test"
	var data []byte = []byte(str)
	action := params["action"]
	if action== "test" {
		w.Write(data)
	}
}
