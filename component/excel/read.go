package excel

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Reader struct {
	file         *excelize.File
	sheetName    string
	headRow      int
	dataStartRow int
	lock         sync.Mutex
}

type ReaderOption struct {
	SheetNumber  int // 0开始
	HeadRow      int // 0开始
	DataStartRow int // 0开始
}

func NewReader(file *excelize.File, option *ReaderOption) *Reader {
	if file == nil || option == nil {
		return nil
	}
	return &Reader{
		file:         file,
		sheetName:    file.GetSheetName(option.SheetNumber),
		headRow:      option.HeadRow,
		dataStartRow: option.DataStartRow,
	}
}

func (r *Reader) Read(dest interface{}) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.headRow >= r.dataStartRow {
		return errors.New("head row must less than data start row")
	}
	rows, getRowsErr := r.file.GetRows(r.sheetName)
	if getRowsErr != nil {
		return nil
	}
	if len(rows) == 0 {
		return errors.New("empty sheet")
	}
	// 需要确保有数据
	if len(rows) <= r.dataStartRow {
		return errors.New("no data")
	}
	// 读取head和data
	headRows := rows[r.headRow]
	dataRows := rows[r.dataStartRow:]

	// 读取head
	headMap := r.getHeadMap(headRows)
	if len(headMap) == 0 {
		return errors.New("empty head")
	}
	// 绑定数据
	if err := r.bindDataToDest(headMap, dataRows, dest); err != nil {
		return err
	}
	// 数据校验
	if err := r.validateData(dest); err != nil {
		return err
	}
	return nil
}

func (r *Reader) getHeadMap(headRows []string) map[string]int {
	headMap := make(map[string]int)
	for i, cell := range headRows {
		if cell == "" {
			continue
		}
		headMap[cell] = i
	}
	return headMap
}

func (r *Reader) bindDataToDest(headMap map[string]int, dataRows [][]string, dest interface{}) error {
	rValue := reflect.ValueOf(dest)
	rType := rValue.Type()
	rKind := rType.Kind()
	if rKind != reflect.Ptr || rValue.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be a pointer to a slice")
	}
	elem := rValue.Type().Elem()
	// if elem.Kind() != reflect.Slice {
	// 	return errors.New("dest structure must be a slice")
	// }
	formatDataList := make([]reflect.Value, 0)
	subElem := elem.Elem()
	for _, dataList := range dataRows {
		if isEmptyLine(dataList) {
			continue
		}
		item := reflect.New(subElem)
		if err := r.bindDataToSt(dataList, headMap, item); err != nil {
			return err
		}
		formatDataList = append(formatDataList, item.Elem())
	}
	rValue.Elem().Set(reflect.Append(rValue.Elem(), formatDataList...))
	return nil
}

func (r *Reader) bindDataToSt(dataList []string, headMap map[string]int, stValue reflect.Value) error {
	elem := stValue.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("[bindDataToSt] elem must be a struct")
	}

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		structField := elem.Type().Field(i)
		if structField.Anonymous {
			return r.bindDataToSt(dataList, headMap, field)
		}
		// 获取名为ex的tag的值
		tagValue := structField.Tag.Get(tagExcel)
		if tagValue == "" {
			continue
		}
		// 获取tag中各个字段的值
		subTagMap := getSubTagMap(tagValue)
		headTag := subTagMap[subTagHead]
		if headTag.param == "" {
			return errors.New("head tag not found")
		}
		subTag, subTagExist := subTagMap[headTag.tag]
		if !subTagExist {
			return errors.New("head not found")
		}
		headIndex, headExist := headMap[subTag.param]
		if !headExist {
			continue
		}
		if headIndex >= len(dataList) {
			continue
		}
		value := Trim(dataList[headIndex])
		// 检查数据类型是否符合预期
		if err := checkFieldTypes(field.Kind(), headTag.param, value); err != nil {
			return err
		}
		if err := setFieldValue(field.Kind(), value, field, headTag.tag); err != nil {
			return err
		}

	}
	return nil
}

func (r *Reader) validateData(data interface{}) error {
	validate := validator.New()
	zh := zh_Hans_CN.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	_ = zhTrans.RegisterDefaultTranslations(validate, trans)
	// 给validate注册一个自定义的标签名称获取函数，使用excel标签值中的head字段名称进行错误提示
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get(tagExcel)
		if tag == "" {
			return fld.Name
		}
		subTagMap := getSubTagMap(tag)
		headTag, headExist := subTagMap[subTagHead]
		if !headExist {
			return fld.Name
		}
		if headTag.param == "" {
			return fld.Name
		}
		return headTag.param
	})
	// 对data进行校验，如果校验不通过，返回错误
	// data必须是一个slice
	rValue := reflect.ValueOf(data).Elem()
	for i := 0; i < rValue.Len(); i++ {
		item := rValue.Index(i).Interface()
		var validationErr validator.ValidationErrors
		if err := validate.Struct(item); err != nil {
			if errors.As(err, &validationErr) {
				for _, v := range validationErr {
					errMsg := v.Translate(trans)
					return fmt.Errorf("rowNumber %d: %s", i, errMsg)
				}
			}
		}
	}
	return nil
}

// checkFieldTypes 检查 Excel 数据类型是否符合预期
func checkFieldTypes(kind reflect.Kind, head, value string) error {
	if head == "" {
		return nil
	}

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		// 如果有逗号，去掉逗号
		newValue := strings.Replace(value, ",", "", -1)
		if _, err := strconv.ParseInt(newValue, 10, 64); err != nil {
			return fmt.Errorf("field %s: expected int64", head)
		}
	case reflect.Float32, reflect.Float64:
		// 如果有逗号，去掉逗号
		newValue := strings.Replace(value, ",", "", -1)
		if _, err := strconv.ParseFloat(newValue, 64); err != nil {
			return fmt.Errorf("field %s: expected float64", head)
		}
	}
	return nil
}

func setFieldValue(kind reflect.Kind, value string, field reflect.Value, key string) error {
	switch kind {
	case reflect.String:
		field.Set(reflect.ValueOf(value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:

		// 如果有逗号，去掉逗号
		value = strings.Replace(value, ",", "", -1)
		uintVal, _ := strconv.ParseUint(value, 10, 64)

		integerKind := kind
		switch integerKind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(uintVal))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uintVal)
		}
	case reflect.Float32, reflect.Float64:
		value = strings.Replace(value, ",", "", -1)
		floatVal, _ := strconv.ParseFloat(value, 64)
		field.SetFloat(floatVal)
	default:
		return fmt.Errorf("field type not support, key: %s", key)
	}
	return nil
}

func isEmptyLine(data []string) bool {
	for _, v := range data {
		if len(v) != 0 {
			return false
		}
	}
	return true
}
