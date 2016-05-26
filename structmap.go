package structmap

import (
	"reflect"
	"strconv"
	"strings"
)

func structToString(v reflect.Value) map[string]interface{} {
	res := make(map[string]interface{}, 1)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("json")
		if key == "" {
			key = v.Type().Field(i).Name
		}
		val := v.Field(i).Interface()
		if val == nil {
			val = nil
		} else if reflect.TypeOf(val).Name() != "string" {
			val = toString(val)
		}
		res[key] = val
	}
	return res
}
func structToStringSlice(v reflect.Value) []map[string]interface{} {
	res := make([]map[string]interface{}, v.Len())
	len := v.Len()
	for i := 0; i < len; i++ {
		res[i] = structToString(v.Index(i))
	}
	return res
}

func structToMap(v reflect.Value) map[string]interface{} {
	res := make(map[string]interface{}, 1)

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(v.Type().Field(i).Name)
		val := v.Field(i).Interface()
		if val == nil {
			val = nil
		}
		res[key] = val
	}
	return res
}

func StructToMap(s interface{}, opts ...bool) map[string]interface{} {
	if len(opts) > 0 && opts[0] {
		return structToString(reflect.ValueOf(s))
	}
	return structToMap(reflect.ValueOf(s))
}

func toString(v interface{}) interface{} {

	if reflect.TypeOf(v).Name() == "int" {
		return strconv.Itoa(v.(int))
	} else if reflect.TypeOf(v).Name() == "int8" {
		return strconv.Itoa(int(v.(int8)))
	} else if reflect.TypeOf(v).Name() == "int16" {
		return strconv.Itoa(int(v.(int16)))
	} else if reflect.TypeOf(v).Name() == "int32" {
		return strconv.Itoa(int(v.(int32)))
	} else if reflect.TypeOf(v).Name() == "int64" {
		return strconv.Itoa(int(v.(int64)))
	} else if reflect.TypeOf(v).Name() == "float32" {
		return strconv.FormatFloat(float64(v.(float32)), 'f', 2, 32)
	} else if reflect.TypeOf(v).Name() == "float64" {
		return strconv.FormatFloat(v.(float64), 'f', 2, 64)
	} else if reflect.TypeOf(v).Name() == "bool" {
		return v.(bool)
	} else {
		if reflect.TypeOf(v).Kind().String() == "slice" {
			return structToStringSlice(reflect.ValueOf(v))
		} else {
			return structToString(reflect.ValueOf(v))
		}
	}

}
