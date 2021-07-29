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

var ss = []Person{
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

func TestStream_Filter(t *testing.T) {

	Ns := streaml.New(ss)
	fmt.Println(ss)
	res := Ns.Filter(`t -> t.Age > 25`).Sort("Name").Min("")

	fmt.Println(res.(Person))
	fmt.Println("ss = ", ss)
	fmt.Println("res = ", res)
	fmt.Println("------------------------")

	s := []int{1, 2, 3, 4, 5, 6}
	TNs := streaml.New(s)
	// if the slice contains the basic type the parameter of Sort and Max or Min methods
	// should be the empty string
	res = TNs.Filter(`a -> a >= 4`).Sort("").Max("")
	fmt.Println(res.(int))
	fmt.Println("------------------------")

	res = TNs.Filter(`c -> c >= 3`).ToSlice()
	fmt.Println(res)

	fmt.Println(TNs.Data)
}

func TestCompare_NewAssign(t *testing.T) {
	testFunc(ss)
}

func testFunc(a interface{}) {
	fmt.Println("Type of a", reflect.TypeOf(a))
	value := reflect.ValueOf(a)
	of := reflect.ValueOf(value.Index(0))
	fmt.Println(of)
	//fmt.Println(comparel.NewAssign(value.Index(0).Interface()).UnsafeAddr())
	//fmt.Println(value.Index(0).UnsafeAddr())
}
