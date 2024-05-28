package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWrite(t *testing.T) {
	type Dest struct {
		SerialNumber int64  `ex:"head:序号,type:int64" validate:"min=10,max=100"`
		UserName     string `ex:"head:姓名"`
		Age          int64  `ex:"head:年龄,type:int64"`
	}
	var dataList []Dest
	dataList = append(dataList, Dest{
		SerialNumber: 1,
		UserName:     "张三",
		Age:          18,
	})
	excelWriter := NewWrite(&WriteOption{
		SheetName: "Sheet1",
		HeadRow:   0,
	})
	buffer, err := excelWriter.GenerateFileStream(dataList)
	assert.Nil(t, err)
	t.Log(buffer.String())
}
