package golang

import (
	"fmt"
	"os"
	"testing"
)

func TestSourceWithoutFunctionBodies(t *testing.T) {
	data, _ := os.ReadFile("bodyless.go")
	stripped, err := sourceWithoutFunctionBodies("bodyless.go", string(data))
	if err != nil {
		t.Fail()
	}
	fmt.Println(stripped)
}
