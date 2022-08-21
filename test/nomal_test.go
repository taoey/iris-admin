package test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func Test01(t *testing.T) {
	file, _ := os.Create("data.txt")
	num := 10000
	for i := 0; i < num; i++ {
		file.WriteString(strconv.Itoa(i))
		file.WriteString("\n")
	}
	file.Close()
}

func TestName(t *testing.T) {
	data := [5]int{1, 2, 3, 4, 5}

	fmt.Println(data[3:4])
}

type people struct {
	Name string
	Age  int
}

type student struct {
	people
	Name string
}

func Test02(t *testing.T) {
	a := student{
		people{Name: "tao", Age: 12}, "s1",
	}
	fmt.Println(a.people.Name)
}
