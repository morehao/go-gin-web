package ctruser

import (
	"go-gin-web/internal/apps/demoapp/dto/dtouser"
	"go-gin-web/internal/apps/demoapp/service/svcuser"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
)

type UserCtr interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Detail(c *gin.Context)
	PageList(c *gin.Context)
}

type userCtr struct {
	userSvc svcuser.UserSvc
}

var _ UserCtr = (*userCtr)(nil)

func NewUserCtr() UserCtr {
	return &userCtr{
		userSvc: svcuser.NewUserSvc(),
	}
}

// Create 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @accept application/json
// @Produce application/json
// @Param req body dtouser.UserCreateReq true "创建用户"
// @Success 200 {object} gincontext.DtoRender{data=dtouser.UserCreateResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /demoapp/user/create [post]
func (ctr *userCtr) Create(c *gin.Context) {
	var req dtouser.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}
	res, err := ctr.userSvc.Create(c, &req)
	if err != nil {
		gincontext.Fail(c, err)
		return
	} else {
		gincontext.Success(c, res)
	}
}

// Delete 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @accept application/json
// @Produce application/json
// @Param req body dtouser.UserDeleteReq true "删除用户"
// @Success 200 {object} gincontext.DtoRender{data=string} "{"code": 0,"data": "ok","msg": "删除成功"}"
// @Router /demoapp/user/delete [post]
func (ctr *userCtr) Delete(c *gin.Context) {
	var req dtouser.UserDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	if err := ctr.userSvc.Delete(c, &req); err != nil {
		gincontext.Fail(c, err)
		return
	} else {
		gincontext.Success(c, "删除成功")
	}
}

// Update 修改用户
// @Tags 用户管理
// @Summary 修改用户
// @accept application/json
// @Produce application/json
// @Param req body dtouser.UserUpdateReq true "修改用户"
// @Success 200 {object} gincontext.DtoRender{data=string} "{"code": 0,"data": "ok","msg": "修改成功"}"
// @Router /demoapp/user/update [post]
func (ctr *userCtr) Update(c *gin.Context) {
	var req dtouser.UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}
	if err := ctr.userSvc.Update(c, &req); err != nil {
		gincontext.Fail(c, err)
		return
	} else {
		gincontext.Success(c, "修改成功")
	}
}

// Detail 用户详情
// @Tags 用户管理
// @Summary 用户详情
// @accept application/json
// @Produce application/json
// @Param req query dtouser.UserDetailReq true "用户详情"
// @Success 200 {object} gincontext.DtoRender{data=dtouser.UserDetailResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /demoapp/user/detail [get]
func (ctr *userCtr) Detail(c *gin.Context) {
	var req dtouser.UserDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}
	res, err := ctr.userSvc.Detail(c, &req)
	if err != nil {
		gincontext.Fail(c, err)
		return
	} else {
		gincontext.Success(c, res)
	}
}

// PageList 用户列表
// @Tags 用户管理
// @Summary 用户列表分页
// @accept application/json
// @Produce application/json
// @Param req query dtouser.UserPageListReq true "用户列表"
// @Success 200 {object} gincontext.DtoRender{data=dtouser.UserPageListResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /demoapp/user/pageList [get]
func (ctr *userCtr) PageList(c *gin.Context) {
	var req dtouser.UserPageListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}
	res, err := ctr.userSvc.PageList(c, &req)
	if err != nil {
		gincontext.Fail(c, err)
		return
	} else {
		gincontext.Success(c, res)
	}
}
