package services

import (
	"reflect"
)

func StructToMapInt(data interface{}) map[string]int {
	result := make(map[string]int)
	elem := reflect.ValueOf(data)

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface().(int)
		result[field] = value
	}

	return result
}
