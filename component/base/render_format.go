package base

import (
	"fmt"
	"go-web/component/utils"
	"reflect"
	"strconv"
)

func ResponseFormat(data interface{}) {
	if data == nil {
		return
	}
	responseFormat(reflect.ValueOf(data))
}

func responseFormat(val reflect.Value) {
	vType := val.Type()
	kd := val.Kind()
	switch kd {
	case reflect.Slice, reflect.Array:
		// tt := vType.Name()
		// vv := val.Interface()
		// fmt.Println(tt, vv)
		fmt.Println("IsNil：", val.IsNil())
		fmt.Println("CanSet：", val.CanSet())
		if val.IsNil() {
			if val.CanSet() {
				newSlice := reflect.MakeSlice(vType, 0, 0)
				val.Set(newSlice)
			}
		} else {
			for i := 0; i < val.Len(); i++ {
				field := val.Index(i)
				responseFormat(field)
			}
		}
	case reflect.Map:
		keys := val.MapKeys()
		for _, v := range keys {
			field := val.MapIndex(v)
			responseFormat(field)
		}
	case reflect.Struct:
		// if val.IsZero() {
		// 	break
		// }
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			// vKind := field.Kind()
			// name := vType.Field(i).Name
			// fmt.Println(name)
			fmt.Println("字段名2：", vType.Field(i).Name)
			responseFormat(field)
		}
	case reflect.Ptr:
		if val.IsZero() {
			break
		}
		st := val.Elem()
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			responseFormat(field)
		}
	case reflect.Float64:

	default:
		return
	}
}

func needFillEmptySlice(kd reflect.Kind) bool {
	switch kd {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct, reflect.Ptr:
		return true
	default:
		return false
	}
}

func isBasicDataType(kd reflect.Kind) bool {
	switch kd {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float64, reflect.Float32,
		reflect.String, reflect.Bool:
		return true
	}

	return false
}

func setFloat64Prec(val reflect.Value, stField reflect.StructField) {
	if val.Kind() == reflect.Float64 {
		precisionTag := stField.Tag.Get("precision")
		if precisionTag != "" {
			precision, _ := strconv.Atoi(precisionTag)
			if precision > 0 {
				newValue := utils.SetPrec(val.Float(), precision)
				val.SetFloat(newValue)
			}
		}
	}
}

// func setFloat64Prec(val reflect.Value) {
// 	if val.Kind() == reflect.Float64 {
// 		precisionTag := val.Type().Field(0).Tag.Get("precision")
// 		if precisionTag != "" {
// 			precision, _ := strconv.Atoi(precisionTag)
// 			if precision > 0 {
// 				originalValue := val.Float()
// 				newValue := utils.SetPrec(originalValue, precision)
// 				val.SetFloat(newValue)
// 			}
// 		}
// 	}
// }

func SetFloat64Prec(data interface{}) {
	v := reflect.ValueOf(data).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("precision")

		if tag != "" {
			prec, _ := strconv.Atoi(tag)

			switch field.Kind() {
			case reflect.Float64:
				field.SetFloat(round(field.Float(), prec))
			case reflect.Slice:
				for j := 0; j < field.Len(); j++ {
					item := field.Index(j)
					if item.Kind() == reflect.Float64 {
						item.SetFloat(round(item.Float(), prec))
					} else if item.Kind() == reflect.Struct {
						SetFloat64Prec(item.Addr().Interface())
					}
				}
			case reflect.Map:
				for _, key := range field.MapKeys() {
					value := field.MapIndex(key)
					if value.Kind() == reflect.Float64 {
						field.SetMapIndex(key, reflect.ValueOf(round(value.Float(), prec)))
					}
				}
			case reflect.Struct:
				SetFloat64Prec(field.Addr().Interface())
			}
		}
	}
}

func round(f float64, prec int) float64 {
	format := "%." + strconv.Itoa(prec) + "f"
	s := fmt.Sprintf(format, f)
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
