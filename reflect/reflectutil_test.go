package reflect

import (
	"testing"
	"fmt"
)

func TestGetType(t *testing.T) {
	getType := GetType(&Info{"test"})
	fmt.Print(getType.Method(0))
}

func TestGetJsonInfo(t *testing.T) {
	GetJsonInfo()
}
