package reflect

import (
	"reflect"
	"fmt"
)

type Info struct {
	Name string
}

func (*Info) InvokeMethod() {
	fmt.Println("test")
}

func GetType() {
	info := &Info{"dream"}

	infotype := reflect.TypeOf(info)
	fmt.Println("string",infotype.Method(0))

	value:=reflect.ValueOf(info)
	fmt.Println("string",value.Method(0))

	for i:=0;i<infotype.NumMethod() ;i++  {
		println(infotype.Method(i).Name)
	}

	methodRes,err:=infotype.MethodByName("InvokeMethod")
	println(err)
	fmt.Println(methodRes)

}
