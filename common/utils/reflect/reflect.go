package utils

import "reflect"

func GetReflectValue(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int:
		return value.Int()
	case reflect.Bool:
		return value.Bool()
	default:
		if value.Len() == 0 {
			return nil
		}
		list := []string{}
		for index := 0; index < value.Len(); index++ {
			list = append(list, value.Index(index).String())
		}
		return list
	}
}
