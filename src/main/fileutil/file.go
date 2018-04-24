package fileutil

import "fmt"

type Reader interface {
	call()
}

func call() {
	fmt.Print("")
}
