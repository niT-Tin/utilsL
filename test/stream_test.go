package stream

import (
	"Utils-Liu/streaml"
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestStream_Filter(t *testing.T) {
	//s := []int{1, 2, 3, 4, 5, 6}

	ss := []Person{
		{"Bruce", 25},
		{"Barry", 23},
		{"Hal", 30},
		{"Dianna", 24},
		{"Ashin", 27},
		{"Max", 21},
		{"Mao", 20},
		{"Test", 99},
		{"Alex", 18},
	}

	Ns := streaml.New(ss)
	fmt.Println(ss)
	res := Ns.Filter(`t -> t.Age > 25`).Sort("Name").Min("")

	fmt.Println(res.(Person))
	fmt.Println("ss = ", ss)
	fmt.Println("res = ", res)
	fmt.Println("------------------------")
}
func testFunc(a interface{}) {
	fmt.Println("Type of a", reflect.TypeOf(a))
	value := reflect.ValueOf(a)
	fmt.Println(reflect.ValueOf(value.Index(0)))
	fmt.Println(value.Index(0))
}
