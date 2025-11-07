package ctruser

import (
	"github.com/morehao/go-gin-web/apps/demoapp/internal/dto/dtouser"
	"github.com/morehao/go-gin-web/apps/demoapp/internal/service/svcuser"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
)

type UserCtr interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Detail(ctx *gin.Context)
	PageList(ctx *gin.Context)
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
// @Tags 用户
// @Summary 创建用户
// @accept application/json
// @Produce application/json
// @Param req body dtouser.UserCreateReq true "创建用户"
// @Success 200 {object} gincontext.DtoRender{data=dtouser.UserCreateResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /demoapp/user/create [post]
func (ctr *userCtr) Create(ctx *gin.Context) {
	var req dtouser.UserCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	res, err := ctr.userSvc.Create(ctx, &req)
	if err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, res)
	}
}

// Delete 删除用户
// @Tags 用户
// @Summary 删除用户
// @accept application/json
// @Produce application/json
// @Param req body dtouser.UserDeleteReq true "删除用户"
// @Success 200 {object} gincontext.DtoRender{data=string} "{"code": 0,"data": "ok","msg": "删除成功"}"
// @Router /demoapp/user/delete [post]
func (ctr *userCtr) Delete(ctx *gin.Context) {
	var req dtouser.UserDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}

	if err := ctr.userSvc.Delete(ctx, &req); err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, "删除成功")
	}
}

// Update 修改用户
// @Tags 用户
// @Summary 修改用户
// @accept application/json
// @Produce application/json
// @Param req body dtouser.UserUpdateReq true "修改用户"
// @Success 200 {object} gincontext.DtoRender{data=string} "{"code": 0,"data": "ok","msg": "修改成功"}"
// @Router /demoapp/user/update [post]
func (ctr *userCtr) Update(ctx *gin.Context) {
	var req dtouser.UserUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	if err := ctr.userSvc.Update(ctx, &req); err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, "修改成功")
	}
}

// Detail 用户详情
// @Tags 用户
// @Summary 用户详情
// @accept application/json
// @Produce application/json
// @Param req query dtouser.UserDetailReq true "用户详情"
// @Success 200 {object} gincontext.DtoRender{data=dtouser.UserDetailResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /demoapp/user/detail [get]
func (ctr *userCtr) Detail(ctx *gin.Context) {
	var req dtouser.UserDetailReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	res, err := ctr.userSvc.Detail(ctx, &req)
	if err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, res)
	}
}

// PageList 用户列表
// @Tags 用户
// @Summary 用户列表分页
// @accept application/json
// @Produce application/json
// @Param req query dtouser.UserPageListReq true "用户列表"
// @Success 200 {object} gincontext.DtoRender{data=dtouser.UserPageListResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /demoapp/user/pageList [get]
func (ctr *userCtr) PageList(ctx *gin.Context) {
	var req dtouser.UserPageListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	res, err := ctr.userSvc.PageList(ctx, &req)
	if err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, res)
	}
}
