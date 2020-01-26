package main

import (
	"fmt"
	"reflect"
)

func reflection(x interface{}, fn func(input string)) {
	val := getValue(x)
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
