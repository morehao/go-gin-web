package svcuser

import (
	"github.com/morehao/go-gin-web/internal/apps/demoapp/code"
	"github.com/morehao/go-gin-web/internal/apps/demoapp/dao/daouser"
	"github.com/morehao/go-gin-web/internal/apps/demoapp/dto/dtouser"
	"github.com/morehao/go-gin-web/internal/apps/demoapp/model"
	"github.com/morehao/go-gin-web/internal/apps/demoapp/object/objcommon"
	"github.com/morehao/go-gin-web/internal/apps/demoapp/object/objuser"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/gutils"
)

type UserSvc interface {
	Create(ctx *gin.Context, req *dtouser.UserCreateReq) (*dtouser.UserCreateResp, error)
	Delete(ctx *gin.Context, req *dtouser.UserDeleteReq) error
	Update(ctx *gin.Context, req *dtouser.UserUpdateReq) error
	Detail(ctx *gin.Context, req *dtouser.UserDetailReq) (*dtouser.UserDetailResp, error)
	PageList(ctx *gin.Context, req *dtouser.UserPageListReq) (*dtouser.UserPageListResp, error)
}

type userSvc struct {
}

var _ UserSvc = (*userSvc)(nil)

func NewUserSvc() UserSvc {
	return &userSvc{}
}

// Create 创建用户管理
func (svc *userSvc) Create(ctx *gin.Context, req *dtouser.UserCreateReq) (*dtouser.UserCreateResp, error) {
	userID := gincontext.GetUserID(ctx)
	insertEntity := &model.UserEntity{
		CompanyID:    req.CompanyID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
		CreatedBy:    userID,
		UpdatedBy:    userID,
	}

	if err := daouser.NewUserDao().Insert(ctx, insertEntity); err != nil {
		glog.Errorf(ctx, "[svcuser.UserCreate] daoUser Create fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, code.GetError(code.UserCreateError)
	}
	return &dtouser.UserCreateResp{
		ID: insertEntity.ID,
	}, nil
}

// Delete 删除用户管理
func (svc *userSvc) Delete(ctx *gin.Context, req *dtouser.UserDeleteReq) error {
	userID := gincontext.GetUserID(ctx)

	if err := daouser.NewUserDao().Delete(ctx, req.ID, userID); err != nil {
		glog.Errorf(ctx, "[svcuser.Delete] daoUser Delete fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return code.GetError(code.UserDeleteError)
	}
	return nil
}

// Update 更新用户管理
func (svc *userSvc) Update(ctx *gin.Context, req *dtouser.UserUpdateReq) error {
	userID := gincontext.GetUserID(ctx)

	updateEntity := &model.UserEntity{
		CompanyID:    req.CompanyID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
		UpdatedBy:    userID,
	}
	if err := daouser.NewUserDao().UpdateByID(ctx, req.ID, updateEntity); err != nil {
		glog.Errorf(ctx, "[svcuser.UserUpdate] daoUser UpdateByID fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return code.GetError(code.UserUpdateError)
	}
	return nil
}

// Detail 根据id获取用户管理
func (svc *userSvc) Detail(ctx *gin.Context, req *dtouser.UserDetailReq) (*dtouser.UserDetailResp, error) {
	detailEntity, err := daouser.NewUserDao().GetById(ctx, req.ID)
	if err != nil {
		glog.Errorf(ctx, "[svcuser.UserDetail] daoUser GetById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, code.GetError(code.UserGetDetailError)
	}
	// 判断是否存在
	if detailEntity == nil || detailEntity.ID == 0 {
		return nil, code.GetError(code.UserNotExistError)
	}
	resp := &dtouser.UserDetailResp{
		ID: detailEntity.ID,
		UserBaseInfo: objuser.UserBaseInfo{
			CompanyID:    detailEntity.CompanyID,
			DepartmentID: detailEntity.DepartmentID,
			Name:         detailEntity.Name,
		},
		OperatorBaseInfo: objcommon.OperatorBaseInfo{
			CreatedBy: detailEntity.CreatedBy,
			CreatedAt: detailEntity.CreatedAt.Unix(),
			UpdatedBy: detailEntity.UpdatedBy,
			UpdatedAt: detailEntity.UpdatedAt.Unix(),
		},
	}
	return resp, nil
}

// PageList 分页获取用户管理列表
func (svc *userSvc) PageList(ctx *gin.Context, req *dtouser.UserPageListReq) (*dtouser.UserPageListResp, error) {
	cond := &daouser.UserCond{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	dataList, total, err := daouser.NewUserDao().GetPageListByCond(ctx, cond)
	if err != nil {
		glog.Errorf(ctx, "[svcuser.UserPageList] daoUser GetPageListByCond fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, code.GetError(code.UserGetPageListError)
	}
	list := make([]dtouser.UserPageListItem, 0, len(dataList))
	for _, v := range dataList {
		list = append(list, dtouser.UserPageListItem{
			ID: v.ID,
			UserBaseInfo: objuser.UserBaseInfo{
				CompanyID:    v.CompanyID,
				DepartmentID: v.DepartmentID,
				Name:         v.Name,
			},
			OperatorBaseInfo: objcommon.OperatorBaseInfo{
				CreatedBy: v.CreatedBy,
				CreatedAt: v.CreatedAt.Unix(),
				UpdatedBy: v.UpdatedBy,
				UpdatedAt: v.UpdatedAt.Unix(),
			},
		})
	}
	return &dtouser.UserPageListResp{
		List:  list,
		Total: total,
	}, nil
}
