package reflect

import (
	"fmt"
	"reflect"
	"io/ioutil"
	"encoding/json"
)

type Info struct {
	Name string
}

func (*Info) InvokeMethod() {
	fmt.Println("test")
}

//获取实例的type
func GetType(interfaceInstance interface{}) reflect.Type {
	typeOf := reflect.TypeOf(interfaceInstance)
	return typeOf
}

type Infos struct {
	Ac []Action
}

type Action struct {
	Path        string
	Handler     string
	Description string
}

func GetJsonInfo() {
	bytes, e := ioutil.ReadFile("/Users/fulei.yang/go/src/gopractice/src/main/resources/mvc.json")
	if e != nil {
		panic(e)
	}

	infos := Infos{}
	json.Unmarshal(bytes, &infos)
	println(infos.Ac[0].Path)

}
