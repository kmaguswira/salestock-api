package utils

import (
	"reflect"
	"strconv"
)

func CopyValue(a, b interface{}) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b).Elem()

	at := av.Type()

	for i := 0; i < at.NumField(); i++ {
		name := at.Field(i).Name
		bf := bv.FieldByName(name)

		if bf.IsValid() && av.Field(i).Type().Name() == "string" && av.Field(i).Interface().(string) != "" {
			if bf.Type().Name() == "string" {
				bf.Set(av.Field(i))
			} else if bf.Type().Name() == "int" {
				newVal, _ := strconv.ParseInt(av.Field(i).Interface().(string), 10, 0)
				bf.SetInt(newVal)
			}

		} else if bf.IsValid() && av.Field(i).Type().Name() == "int" && bf.Type().Name() == "int" && bf.Interface().(int) != av.Field(i).Interface().(int) {
			bf.Set(av.Field(i))

		}
	}

}
