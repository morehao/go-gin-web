package ctrUser

import (
	"go-gin-web/internal/app/dto/dtoUser"
	"go-gin-web/internal/app/service/svcUser"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gcontext/ginrender"
)

type UserCtr interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Detail(c *gin.Context)
	PageList(c *gin.Context)
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

// Create 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @accept application/json
// @Produce application/json
// @Param req body dtoUser.UserCreateReq true "创建用户"
// @Success 200 {object} dto.DefaultRender{data=dtoUser.UserCreateResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /app/user/create [post]
func (ctr *userCtr) Create(c *gin.Context) {
	var req dtoUser.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ginrender.Fail(c, err)
		return
	}
	res, err := ctr.userSvc.Create(c, &req)
	if err != nil {
		ginrender.Fail(c, err)
		return
	} else {
		ginrender.Success(c, res)
	}
}

// Delete 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @accept application/json
// @Produce application/json
// @Param req body dtoUser.UserDeleteReq true "删除用户"
// @Success 200 {object} dto.DefaultRender{data=string} "{"code": 0,"data": "ok","msg": "删除成功"}"
// @Router /app/user/delete [post]
func (ctr *userCtr) Delete(c *gin.Context) {
	var req dtoUser.UserDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ginrender.Fail(c, err)
		return
	}

	if err := ctr.userSvc.Delete(c, &req); err != nil {
		ginrender.Fail(c, err)
		return
	} else {
		ginrender.Success(c, "删除成功")
	}
}

// Update 修改用户
// @Tags 用户管理
// @Summary 修改用户
// @accept application/json
// @Produce application/json
// @Param req body dtoUser.UserUpdateReq true "修改用户"
// @Success 200 {object} dto.DefaultRender{data=string} "{"code": 0,"data": "ok","msg": "修改成功"}"
// @Router /app/user/update [post]
func (ctr *userCtr) Update(c *gin.Context) {
	var req dtoUser.UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ginrender.Fail(c, err)
		return
	}
	if err := ctr.userSvc.Update(c, &req); err != nil {
		ginrender.Fail(c, err)
		return
	} else {
		ginrender.Success(c, "修改成功")
	}
}

// Detail 用户详情
// @Tags 用户管理
// @Summary 用户详情
// @accept application/json
// @Produce application/json
// @Param req query dtoUser.UserDetailReq true "用户详情"
// @Success 200 {object} dto.DefaultRender{data=dtoUser.UserDetailResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /app/user/detail [get]
func (ctr *userCtr) Detail(c *gin.Context) {
	var req dtoUser.UserDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginrender.Fail(c, err)
		return
	}
	res, err := ctr.userSvc.Detail(c, &req)
	if err != nil {
		ginrender.Fail(c, err)
		return
	} else {
		ginrender.Success(c, res)
	}
}

// PageList 用户列表
// @Tags 用户管理
// @Summary 用户列表分页
// @accept application/json
// @Produce application/json
// @Param req query dtoUser.UserPageListReq true "用户列表"
// @Success 200 {object} dto.DefaultRender{data=dtoUser.UserPageListResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /app/user/pageList [get]
func (ctr *userCtr) PageList(c *gin.Context) {
	var req dtoUser.UserPageListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ginrender.Fail(c, err)
		return
	}
	res, err := ctr.userSvc.PageList(c, &req)
	if err != nil {
		ginrender.Fail(c, err)
		return
	} else {
		ginrender.Success(c, res)
	}
}
