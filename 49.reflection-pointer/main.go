package main

import "reflect"


func walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	// value 有一个方法 NumField，它返回值中的字段数。这让我们遍历字段并调用 fn 通过我们的测试。
	for i:=0; i<val.NumField();i++ {

		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}