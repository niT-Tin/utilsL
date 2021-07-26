package sortl

import (
	"Utils-Liu/comparel"
	"reflect"
)

// SelectionSort 使用空接口指针满足更多类型
// field 为要比较的具体字段
// dst为一个要比较类型的数据元素
func SelectionSort(arr interface{}, field string) {
	// 通过反射获得interface{}是否为slice类型，如果是则进行后续操作
	if slice := reflect.ValueOf(arr); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			minI := i
			for j := i; j < slice.Len(); j++ {
				if comparel.DeepLesser(slice.Index(j).Interface(), slice.Index(minI).Interface(), field) {
					minI = j
				}
			}
			comparel.DeepSwap(slice.Index(minI), slice.Index(i))
		}
	}
}
