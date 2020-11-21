package entity

import (
	"fmt"
	"testing"
)

func TestOKResponse(t *testing.T) {
	student := Student{"tao"}
	response := OKResponse(student)
	fmt.Println(response)
}

type Student struct {
	name string
}
