package excel

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Reader struct {
	file         *excelize.File
	sheetName    string
	headerRow    int
	dataStartRow int
	lock         sync.Mutex
}

type ReaderOption struct {
	SheetNumber  int // 0开始
	HeaderRow    int // 0开始
	DataStartRow int // 0开始
}

func NewReader(file *excelize.File, option *ReaderOption) *Reader {
	if file == nil || option == nil {
		return nil
	}
	return &Reader{
		file:         file,
		sheetName:    file.GetSheetName(option.SheetNumber),
		headerRow:    option.HeaderRow,
		dataStartRow: option.DataStartRow,
	}
}

func (r *Reader) Read(dest interface{}) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.headerRow >= r.dataStartRow {
		return errors.New("header row must less than data start row")
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
	// 读取header和data
	headerRows := rows[r.headerRow]
	dataRows := rows[r.dataStartRow:]

	// 读取header
	headerMap := r.getHeaderMap(headerRows)
	if len(headerMap) == 0 {
		return errors.New("empty header")
	}
	// 绑定数据
	if err := r.bindDataToDest(headerMap, dataRows, dest); err != nil {
		return err
	}
	return nil
}

func (r *Reader) getHeaderMap(headerRows []string) map[string]int {
	headerMap := make(map[string]int)
	for i, cell := range headerRows {
		if cell == "" {
			continue
		}
		headerMap[cell] = i
	}
	return headerMap
}

func (r *Reader) bindDataToDest(headerMap map[string]int, dataRows [][]string, dest interface{}) error {
	rValue := reflect.ValueOf(dest)
	rType := rValue.Type()
	rKind := rType.Kind()
	if rKind != reflect.Ptr {
		return errors.New("dest must be a pointer")
	}
	elem := rValue.Type().Elem()
	if elem.Kind() != reflect.Slice {
		return errors.New("dest structure must be a slice")
	}
	formatDataList := make([]reflect.Value, 0)
	subElem := elem.Elem()
	for _, dataList := range dataRows {
		if isEmptyLine(dataList) {
			continue
		}
		item := reflect.New(subElem)
		if err := r.bindDataToSt(dataList, headerMap, item); err != nil {
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
		fieldType := elem.Type().Field(i)
		if fieldType.Anonymous {
			return r.bindDataToSt(dataList, headMap, field)
		}
		// 获取名为ex的tag的值
		tagValue := fieldType.Tag.Get(TagExcel)
		if tagValue == "" {
			continue
		}
		// 获取tag中各个字段的值
		subTagMap := getSubTagMap(tagValue)
		headTag := subTagMap[SubTagHeader]
		if headTag.tag == "" {
			continue
		}
		// TODO: 获取对应的参数有问题
		headIndex, ok := headMap[headTag.tag]
		if !ok {
			return errors.New("header not found")
		}
		if headIndex >= len(dataList) {
			continue
		}
		value := Trim(dataList[headIndex])
		if err := setFieldValue(field.Kind(), value, field, headTag.tag); err != nil {
			return err
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
		field.Set(reflect.ValueOf(uintVal))
	case reflect.Float32, reflect.Float64:
		value = strings.Replace(value, ",", "", -1)
		v, _ := strconv.ParseFloat(value, 64)
		field.Set(reflect.ValueOf(v))
	default:
		return errors.New(fmt.Sprintf("field type not support, key: %s", key))
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
