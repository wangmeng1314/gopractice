package httputil

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"runtime"
	"strings"
	"strconv"
	"time"
	"io/ioutil"
)



func HandleRequest() {
	rounter := mux.NewRouter()
	rounter.HandleFunc("/info/{action}", HandleGet).Methods("GET")
	rounter.Handle("/handler/{action}",&Example{})
	http.ListenAndServe(":8080", rounter)
}

func HandleGet(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	action := params["action"]
	go ServiceInfo()
	time.Sleep(time.Second)
	if action == "test" {
		res, err:=ioutil.ReadFile("info.json")
		if nil != err {
			panic(err)
		}
		w.Write(res)
	}
}

func ServiceInfo() interface{} {
	fmt.Println("test")
	var grountieId = GoID()
	fmt.Println(grountieId)
	fmt.Println("start sleep")
	time.Sleep(time.Second)
	fmt.Println("end sleep")
	return 1000
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
