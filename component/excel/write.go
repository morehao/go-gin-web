package excel

import (
	"fmt"
	"reflect"
	"sync"
)

type Write struct {
	sheetName string
	headRow   int
	lock      sync.Mutex
}

type WriteOption struct {
	SheetName string // 表名
	HeadRow   int    // 0开始
}

func NewWrite(option *WriteOption) *Write {
	if option == nil {
		return nil
	}
	return &Write{
		sheetName: option.SheetName,
		headRow:   option.HeadRow,
	}
}

func (w *Write) Write(data interface{}) error {
	w.lock.Lock()
	defer w.lock.Unlock()

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}

	// 检查切片是否为空
	if dataValue.Len() == 0 {
		// 创建一个新实例以获取头行数据
		elemType := dataValue.Type().Elem()
		tempInstance := reflect.New(elemType).Elem().Interface()
		headRows := w.getHeadRows(tempInstance)
		fmt.Println(headRows)
		return nil
	}

	// 在循环外获取一次头行数据
	headRows := w.getHeadRows(dataValue.Index(0).Interface())
	fmt.Println(headRows)

	// 遍历切片元素并处理数据
	for i := 0; i < dataValue.Len(); i++ {
		elementValue := dataValue.Index(i)
		if elementValue.Kind() != reflect.Struct {
			return fmt.Errorf("data must be a slice of structs")
		}
		// 处理数据的逻辑可以放在这里
	}

	return nil
}

func (w *Write) getHeadRows(elem interface{}) []string {
	var headRows []string
	elemValue := reflect.ValueOf(elem)
	elemType := elemValue.Type()
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		tagValue := field.Tag.Get(tagExcel)
		if tagValue == "" {
			continue
		}
		headTag := getHeadTag(tagValue)
		if headTag == nil || headTag.param == "" {
			continue
		}
		headRows = append(headRows, headTag.param)
	}
	return headRows
}
