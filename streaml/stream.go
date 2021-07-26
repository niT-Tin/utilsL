package streaml

import (
	utils "Utils-Liu/comparel"
	"Utils-Liu/sortl"
	"errors"
	"reflect"
	"strings"
	"unsafe"
)

// temp storage for stream data

var fieldL, operatorL string

//var t reflect.Type
const (
	operator    = "->"
	errMsg      = "invalid lambda expression"
	typeMsg     = "data must be slice or map"
	fieldMsg    = "data doesn't have "
	operatorMsg = "operator error"
)

type Stream struct {
	Data       interface{}
	types      reflect.Type
	singleType interface{}
	temp       []interface{}
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

func New(d interface{}, s interface{}) *Stream {
	var types = make([]interface{}, 0)
	var t1 = Stream{
		d,
		reflect.TypeOf(d),
		s,
		types,
	}
	//t = reflect.TypeOf(d)
	return &t1
}

func isSlice(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Slice
}

func isMap(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Map
}

func (s *Stream) Filter(lambda string) *Stream {
	// s := []int{1, 2, 3, 4, 5, 6}
	// Ns := utils.New(s)
	// res := Ns.Filter(`t -> t >= 3`)
	// res should be the [3, 4, 5, 6]
	// separate the lambda expression
	obj, fields, arg, operator := lambdaSplit(lambda)
	fieldL = fields[0]
	operatorL = operator
	if strings.Compare(obj[0], obj[1]) != 0 {
		panic(errMsg)
	}
	s.filter(fields[0], arg, operator)
	return s
}

func (s *Stream) filter(field, arg, operator string) *Stream {
	if !isSlice(s.Data) {
		panic(typeMsg)
	}
	// 检查字段名是否有效
	value := getInternalValue(s.Data)
	if value.FieldByName(field).IsZero() {
		panic(fieldMsg)
	}
	return s.Operate(field, arg, operator)
}

func getInternalValue(a interface{}) reflect.Value {
	if s := reflect.TypeOf(a); s.Kind() == reflect.Slice {
		return reflect.ValueOf(reflect.ValueOf(a).Index(0).Interface())
	}
	return reflect.Value{}
}

func noForOp(f func(a, b interface{}, field string) bool,
	s *Stream, value reflect.Value, field string) {
	var b bool
	for i := 0; i < reflect.ValueOf(s.Data).Len(); i++ {
		if b = f(reflect.ValueOf(s.Data).Index(i).Interface(),
			value.Elem().Interface(), field); b {
			s.temp = append(s.temp, reflect.ValueOf(s.Data).Index(i).Interface())
		}
	}
}

func noForMM(f func(a, b interface{}, field string) bool, s *Stream, fd string) interface{} {
	m := s.temp[0]
	for i := 0; i < len(s.temp); i++ {
		if f(m, s.temp[i], fd) {
			m = s.temp[i]
		}
	}
	return m
}

func (s *Stream) Operate(field, arg, operator string) *Stream {
	value := reflect.New(reflect.TypeOf(s.singleType))
	utils.SwitchTypeSetValue(value.Elem().FieldByName(field), arg)
	switch operator {
	case ">":
		noForOp(utils.DeepGreater, s, value, field)
	case "<":
		noForOp(utils.DeepLesser, s, value, field)
	case ">=":
		noForOp(utils.DeepGEqual, s, value, field)
	case "<=":
		noForOp(utils.DeepLEqual, s, value, field)
	case "==":
		noForOp(utils.DeepEqual, s, value, field)
	default:
		panic(operatorMsg)
	}
	return s
}

func checkLambdaValid(lambda string) error {
	var e error
	if !strings.Contains(lambda, operator) {
		e = errors.New(errMsg + ": " + lambda)
	}
	n := strings.Split(lambda, operator)
	if strings.Compare(n[0], " ") == 0 || strings.Compare(n[0], "") == 0 {
		e = errors.New(errMsg)
	}
	return e
}

func (s *Stream) Sort(sortField string) *Stream {
	if sortField == "" {
		sortl.SelectionSort(s.temp, fieldL)
	} else {
		sortl.SelectionSort(s.temp, sortField)
	}
	return s
}

func max(s *Stream, field string) interface{} {
	if field == "" {
		return noForMM(utils.DeepLesser, s, fieldL)
	} else {
		return noForMM(utils.DeepLesser, s, field)
	}
}

func min(s *Stream, field string) interface{} {
	if field == "" {
		return noForMM(utils.DeepGreater, s, fieldL)
	} else {
		return noForMM(utils.DeepGreater, s, field)
	}
}

func (s *Stream) Count() int {
	return len(s.temp)
}

func (s *Stream) Max(sortField string) interface{} {
	return max(s, sortField)
}

func (s *Stream) Min(sortField string) interface{} {
	return min(s, sortField)
}

func (s *Stream) ToSlice() []interface{} {
	tempSlice := make([]interface{}, 0)
	for i := 0; i < len(s.temp); i++ {
		tempSlice = append(tempSlice, s.temp[i])
	}
	return tempSlice
}

func (s *Stream) toMap() map[string]interface{} {
	return map[string]interface{}{}
}

// 将输入的lambda表达式进行切割
func lambdaSplit(lambda string) (obj []string, fields []string,
	arg string, operator string) {
	obj = make([]string, 2)
	fields = make([]string, 1)
	if err := checkLambdaValid(lambda); err != nil {
		panic(err)
	}
	field := strings.Fields(lambda)
	split := strings.Split(field[2], ".")
	obj[0] = field[0]
	obj[1] = split[0]
	fields[0] = split[1]
	arg = field[len(field)-1]
	operator = field[len(field)-2]
	return obj, fields, arg, operator
}
