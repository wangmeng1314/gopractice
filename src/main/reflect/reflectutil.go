package reflect

import (
	"reflect"
	"fmt"
)

type Info struct {
	Name string
}

func (i *Info) InvokeMethod() {
	fmt.Println("test")
}

func GetType() {
	info := Info{"dream"}

	infotype := reflect.TypeOf(info)
	println(infotype.Name())
	for i:=0;i<infotype.NumMethod() ;i++  {
		println(infotype.Method(i).Name)
	}
	methodRes,err:=infotype.MethodByName("InvokeMethod")
	println(err)
	fmt.Println(methodRes)
}
