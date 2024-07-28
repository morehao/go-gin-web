package svcExample

import (
	"go-gin-web/internal/app/dto/dtoExample"

	"github.com/gin-gonic/gin"
)

type ExampleSvc interface {
	FormatData(c *gin.Context) *dtoExample.FormatDataRes
}

var _ ExampleSvc = (*exampleSvc)(nil)

type exampleSvc struct {
}

func NewExampleSvc() ExampleSvc {
	return &exampleSvc{}
}

func (svc *exampleSvc) FormatData(c *gin.Context) *dtoExample.FormatDataRes {
	return &dtoExample.FormatDataRes{
		Items: []dtoExample.Item{
			{
				Price:     1.22245,
				PriceList: []float64{1.22245, 1.22255},
			},
		},
		Item: dtoExample.Item{
			Price:     1.22245,
			PriceList: []float64{1.22245, 1.22255},
		},
		ItemMap: map[string]dtoExample.Item{
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
