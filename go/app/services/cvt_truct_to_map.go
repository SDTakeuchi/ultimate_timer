package services

import (
	"reflect"
)

	func StructToMap(data interface{}) map[string]int {
		result := make(map[string]int)
		elem := reflect.ValueOf(data).Elem()
		size := elem.NumField()
	
		for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface().(int)
		result[field] = value
		}
	
		return result
	}
