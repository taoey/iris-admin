package resp

import (
	"fmt"
	"testing"
)

func TestOKResponse(t *testing.T) {
	student := Student{"tao"}
	response := OkResponseWithRet(student)
	fmt.Println(response)
}

type Student struct {
	name string
}
