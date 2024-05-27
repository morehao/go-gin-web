package excel

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestReader(t *testing.T) {
	f, err := excelize.OpenFile("read.xlsx")
	assert.Nil(t, err)
	type Dest struct {
		SerialNumber int64  `ex:"head:序号,type:int64" validate:"min=10,max=100"`
		UserName     string `ex:"head:姓名"`
		Age          int64  `ex:"head:年龄,type:int64"`
	}
	var dataList []Dest
	excelReader := NewReader(f, &ReaderOption{
		SheetNumber:  0,
		HeadRow:      0,
		DataStartRow: 1,
	})
	readerErr := excelReader.Read(&dataList)
	assert.Nil(t, readerErr)
	res, _ := jsoniter.MarshalToString(dataList)
	fmt.Println(res)
}
