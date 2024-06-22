package svcUser

import (
	"go-gin-web/internal/demo/dto/dtoUser"

	"github.com/gin-gonic/gin"
)

type UserSvc interface {
	Get(c *gin.Context, req *dtoUser.GetUserReq) (*dtoUser.GetUserRes, error)
	FormatData(c *gin.Context) *dtoUser.FormatDataRes
}

var _ UserSvc = (*userSvc)(nil)

type userSvc struct {
}

func NewUserSvc() UserSvc {
	return &userSvc{}
}

func (svc *userSvc) Get(c *gin.Context, req *dtoUser.GetUserReq) (*dtoUser.GetUserRes, error) {
	return &dtoUser.GetUserRes{}, nil
}

func (svc *userSvc) FormatData(c *gin.Context) *dtoUser.FormatDataRes {
	return &dtoUser.FormatDataRes{
		Items: []dtoUser.Item{
			{
				Price:     1.22245,
				PriceList: []float64{1.22245, 1.22255},
			},
		},
		Item: dtoUser.Item{
			Price:     1.22245,
			PriceList: []float64{1.22245, 1.22255},
		},
		ItemMap: map[string]dtoUser.Item{
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
