package ctrUser

import (
	"go-gin-web/internal/demo/dto/dtoUser"
	"go-gin-web/internal/demo/service/svcUser"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gcore/ginRender"
)

type UserCtr interface {
	Get(c *gin.Context)
	FormatData(c *gin.Context)
}

type userCtr struct {
	userSvc svcUser.UserSvc
}

var _ UserCtr = (*userCtr)(nil)

func NewUserCtr() UserCtr {
	return &userCtr{
		userSvc: svcUser.NewUserSvc(),
	}
}

// Get 获取用户详情
// @Tags 用户管理
// @Summary 获取用户详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param x-token header string true "token"
// @Param req query dtoUser.GetUserReq true "获取用户详情"
// @Success 200 {object} dto.DefaultRender(data=dtoUser.GetUserRes) "{"code": 0,"data": "ok","msg": "success"}"
// @Router /user/get [post]
func (ctr *userCtr) Get(c *gin.Context) {
	var req dtoUser.GetUserReq
	if err := c.ShouldBind(&req); err != nil {
		ginRender.RenderFail(c, err)
		return
	}
	res, err := ctr.userSvc.Get(c, &req)
	if err != nil {
		ginRender.RenderFail(c, err)
		return
	}
	ginRender.RenderSuccess(c, res)
}

func (ctr *userCtr) FormatData(c *gin.Context) {
	res := ctr.userSvc.FormatData(c)

	ginRender.RenderSuccessWithFormat(c, res)
}
