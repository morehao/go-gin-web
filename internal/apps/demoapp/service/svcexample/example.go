package svcexample

import (
	"go-gin-web/internal/apps/demoapp/dto/dtoexample"

	"github.com/gin-gonic/gin"
)

type ExampleSvc interface {
	FormatData(c *gin.Context) *dtoexample.FormatDataRes
}

var _ ExampleSvc = (*exampleSvc)(nil)

type exampleSvc struct {
}

func NewExampleSvc() ExampleSvc {
	return &exampleSvc{}
}

func (svc *exampleSvc) FormatData(c *gin.Context) *dtoexample.FormatDataRes {
	return &dtoexample.FormatDataRes{
		Items: []dtoexample.Item{
			{
				Price:     1.22245,
				PriceList: []float64{1.22245, 1.22255},
			},
		},
		Item: dtoexample.Item{
			Price:     1.22245,
			PriceList: []float64{1.22245, 1.22255},
		},
		ItemMap: map[string]dtoexample.Item{
			"1": {
				Price:     1.22245,
				PriceList: []float64{1.22245, 1.22255},
			},
			"2": {
				Price: 1.22245,
			},
		},
		NameMap: map[string][]string{
			"a": []string{},
		},
		PriceList: []float64{1.22245, 1.22255},
		PriceMap: map[string]float64{
			"1": 1.22245,
		},
	}
}
