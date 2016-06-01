package structmap

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
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
		} else if key == "-" {
			continue
		}

		val := v.Field(i).Interface()
		if val == nil {
			val = nil
		} else if reflect.TypeOf(val).Name() == "Time" {
			format := v.Type().Field(i).Tag.Get("time_format")
			val = timeToString(val, format)
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
	log.Println(reflect.TypeOf(v).Name())

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

func timeToString(v interface{}, format string) string {
	if v.(time.Time).IsZero() {
		return ""
	}

	if format != "" {
		return convertDateTime(v.(time.Time), format)
	}
	return v.(time.Time).String()
}

// ConvertDateTime not yet implemented, need to see the performance first
func convertDateTime(t time.Time, form string) string {

	//map for parsing
	r := strings.NewReplacer(
		"YYYY", "2006",
		"MMMM", "January",
		"MMM", "Jan",
		"MM", "01",
		"M", "1",
		"DDDD", "Mon",
		"DD", "02",
		"D", "2",
		"HH24", "15",
		"HH", "03",
		"H", "3",
		"NN", "04",
		"N", "4",
		"SS", "05",
		"S", "5",
		"AMPM", "PM",
		"ampm", "pm",
	)

	return t.Format(r.Replace(form))
}
