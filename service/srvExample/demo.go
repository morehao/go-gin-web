package srvExample

import (
	"github.com/gin-gonic/gin"
	"go-web/dto/dtoExample"
)

func FormatData(ctx *gin.Context) *dtoExample.FormatDataRes {
	return &dtoExample.FormatDataRes{
		Items: []dtoExample.Item{
			{
				Price: 1.22245,
				// PriceList: []float64{1.22245, 1.22255},
				// NameList:  []string{},
			},
		},
		Item: dtoExample.Item{
			Price: 1.22245,
			// PriceList: []float64{1.22245, 1.22255},
		},
		ItemMap: map[string]dtoExample.Item{
			"1": {
				Price: 1.22245,
				// PriceList: []float64{1.22245, 1.22255},
			},
		},
		NameMap: map[string][]string{
			"a": []string{},
		},
		// PriceList: []float64{1.22245, 1.22255},
		PriceMap: map[string]float64{
			"1": 1.22245,
		},
	}
}
