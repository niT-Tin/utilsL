package utils

import (
	"reflect"
	"strconv"
	"strings"
)

// DeepGreater 深度比较a的field字段的值是否大于b的field字段的值
func DeepGreater(a, b interface{}, field string) bool {
	return deep(a, b, field, deepValueGreater)
}

// DeepLesser 深度比较a的field字段的值是否小于b的field字段的值
func DeepLesser(a, b interface{}, field string) bool {
	return deep(a, b, field, deepValueLesser)
}

// DeepGEqual 深度比较a的field字段的值是否大于小于b的field字段的值
func DeepGEqual(a, b interface{}, field string) bool {
	return deep(a, b, field, deepValueGEqual)
}

// DeepLEqual 深度比较a的field字段的值是否小于等于b的field字段的值
func DeepLEqual(a, b interface{}, field string) bool {
	return deep(a, b, field, deepValueLEqual)
}

// DeepEqual 深度比较a的field字段是否等于b的field字段的值
func DeepEqual(a, b interface{}, field string) bool {
	return deep(a, b, field, deepValueEqual)
}

// DeepSwap 交换a和b两个值
func DeepSwap(a, b reflect.Value) {
	temp := reflect.New(reflect.TypeOf(a.Interface()))
	typeOf := reflect.TypeOf(a.Interface())
	v1 := reflect.ValueOf(a.Interface())
	v2 := reflect.ValueOf(b.Interface())
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		temp.Elem().Field(i).Set( /*fmt.Println(*/ v1.Field(i) /*)*/)
	}
	a.Set(v2)
	b.Set(temp.Elem())
}

// 使用指定函数进行深度比较
func deep(a interface{}, b interface{}, field string, f func(float64) bool) bool {
	name1, name2, typ, b2, done := beforeCompare(a, b, field)
	if done {
		return b2
	}
	return deepCompare(typ, name1, name2, f)
}

// 在深度比较前获取a与b的Type以及Value
func beforeCompare(a interface{}, b interface{}, field string) (reflect.Value, reflect.Value, reflect.Kind, bool, bool) {
	if a == nil || b == nil {
		return reflect.Value{}, reflect.Value{}, 0, a == b, true
	}
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)
	if v1.Type() != v2.Type() {
		return reflect.Value{}, reflect.Value{}, 0, false, true
	}
	var name1, name2 reflect.Value
	if field != "" {
		name1 = v1.FieldByName(field)
		name2 = v2.FieldByName(field)
	} else {
		name1 = v1
		name2 = v2
	}
	typ := name1.Kind()
	return name1, name2, typ, false, false
}

func SwitchTypeSetValue(k reflect.Value, v string) {
	switch k.Kind() {
	case reflect.String:
		k.SetString(v)
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		atoi, _ := strconv.Atoi(v)
		k.SetInt(int64(atoi))
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		atoi, _ := strconv.ParseUint(v, 10, 64)
		k.SetUint(atoi)
	case reflect.Float32, reflect.Float64:
		float, _ := strconv.ParseFloat(v, 64)
		k.SetFloat(float)
	}
}

// 获取指定字段的类型，并调用指定函数进行比较
func deepCompare(typ reflect.Kind, name1, name2 reflect.Value, f func(float64) bool) bool {
	switch typ {
	case reflect.String:
		return f(float64(strings.Compare(name1.String(), name2.String())))
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return f(float64(name1.Int() - name2.Int()))
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return f(float64(name1.Uint() - name2.Uint()))
	case reflect.Float32, reflect.Float64:
		return f(name1.Float() - name2.Float())
	default:
		return false
	}
}

func deepValueEqual(v1 float64) bool {
	if v1 == 0 {
		return true
	} else {
		return false
	}
}

func deepValueGreater(v1 float64) bool {
	if v1 > 0 {
		return true
	} else {
		return false
	}
}

func deepValueLesser(v1 float64) bool {
	if v1 < 0 {
		return true
	} else {
		return false
	}
}

func deepValueLEqual(v1 float64) bool {
	if v1 <= 0 {
		return true
	} else {
		return false
	}
}

func deepValueGEqual(v1 float64) bool {
	if v1 >= 0 {
		return true
	} else {
		return false
	}
}
