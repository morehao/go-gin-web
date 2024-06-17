package base

import (
	"fmt"
	"math"
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
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	vType := val.Type()
	kd := val.Kind()

	switch kd {
	case reflect.Slice, reflect.Array:
		if val.IsNil() {
			if val.CanSet() {
				newSlice := reflect.MakeSlice(vType, 0, 0)
				val.Set(newSlice)
			}
		} else {
			for i := 0; i < val.Len(); i++ {
				field := val.Index(i)
				if field.Kind() == reflect.Float64 {
					// TODO：格式化精度
				} else {
					responseFormat(field)
				}
			}
		}
	case reflect.Map:
		mapRange := val.MapRange()
		for mapRange.Next() {
			key := mapRange.Key()
			value := mapRange.Value()
			switch value.Kind() {
			case reflect.Ptr, reflect.Interface:
				if !value.IsNil() {
					responseFormat(value.Elem())
				}
			case reflect.Struct:
				newValue := reflect.New(value.Type()).Elem()
				newValue.Set(value)
				responseFormat(newValue.Addr())
				val.SetMapIndex(key, newValue)
			case reflect.Slice, reflect.Array:
				if value.IsNil() {
					if value.CanSet() {
						newSlice := reflect.MakeSlice(vType, 0, 0)
						value.Set(newSlice)
					}
				} else {
					for i := 0; i < value.Len(); i++ {
						field := value.Index(i)
						responseFormat(field)
					}
				}
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			typeField := vType.Field(i)
			switch field.Kind() {
			case reflect.Ptr, reflect.Struct, reflect.Interface:
				responseFormat(field)
			case reflect.Map:
				elem := field.Type().Elem()
				if elem.Kind() == reflect.Float64 {
					// 获取精度
					precisionTag := typeField.Tag.Get("precision")
					if precisionTag != "" && field.CanSet() {
						precision, _ := strconv.Atoi(precisionTag)
						// 遍历map
						mapRange := field.MapRange()
						for mapRange.Next() {
							mapKey := mapRange.Key()
							mapValue := mapRange.Value()
							if mapValue.Kind() == reflect.Float64 {
								newValue := reflect.New(mapValue.Type()).Elem()
								// 给mapValue赋值，新值经过精度处理
								newValue.SetFloat(round(mapValue.Float(), precision))
								// 给map赋值
								field.SetMapIndex(mapKey, newValue)
							}
						}
					}

				} else {
					responseFormat(field)
				}

			case reflect.Float64:
				setFieldPrecision(field, typeField)
			case reflect.Slice, reflect.Array:
				if field.IsNil() {
					if field.CanSet() {
						newSlice := reflect.MakeSlice(field.Type(), 0, 0)
						field.Set(newSlice)
					}
				} else {
					for j := 0; j < field.Len(); j++ {
						subField := field.Index(j)
						if subField.Kind() == reflect.Float64 {
							setFieldPrecision(subField, typeField)
						} else {
							responseFormat(subField)
						}
					}
				}

			}
			if field.Kind() == reflect.Float64 {
				setFieldPrecision(field, typeField)
			}

		}
	case reflect.Ptr:
		if !val.IsNil() {
			st := val.Elem()
			for i := 0; i < st.NumField(); i++ {
				field := st.Field(i)
				typeField := st.Type().Field(i)

				if field.Kind() == reflect.Float64 {
					setFieldPrecision(field, typeField)
				}

				responseFormat(field)
			}
		}
	}
}

func setFieldPrecision(field reflect.Value, typeField reflect.StructField) {
	precisionTag := typeField.Tag.Get("precision")
	if precisionTag != "" && field.CanSet() {
		precision, err := strconv.Atoi(precisionTag)
		if err != nil {
			fmt.Println("Invalid precision:", err)
			return
		}
		rounded := round(field.Float(), precision)
		field.SetFloat(rounded)
	}
}

func round(x float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(x*pow) / pow
}
