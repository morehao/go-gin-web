package excel

import (
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestReader(t *testing.T) {
	f, err := excelize.OpenFile("read.xlsx")
	assert.Nil(t, err)
	type Dest struct {
		SerialNumber int64  `ex:"header:序号;type:int64"`
		UserName     string `ex:"header:用户名;type:string"`
		Age          int64  `ex:"header:年龄;type:int64"`
	}
	var DestList []Dest
	excelReader := NewReader(f, &ReaderOption{
		SheetNumber:  0,
		HeaderRow:    0,
		DataStartRow: 1,
	})
	readerErr := excelReader.Read(&DestList)
	assert.Nil(t, readerErr)
}
