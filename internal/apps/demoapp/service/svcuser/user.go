package svcuser

import (
	"go-gin-web/internal/apps/demoapp/code"
	"go-gin-web/internal/apps/demoapp/dao/daouser"
	"go-gin-web/internal/apps/demoapp/dto/dtouser"
	"go-gin-web/internal/apps/demoapp/model"
	"go-gin-web/internal/apps/demoapp/object/objcommon"
	"go-gin-web/internal/apps/demoapp/object/objuser"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/gutils"
)

type UserSvc interface {
	Create(c *gin.Context, req *dtouser.UserCreateReq) (*dtouser.UserCreateResp, error)
	Delete(c *gin.Context, req *dtouser.UserDeleteReq) error
	Update(c *gin.Context, req *dtouser.UserUpdateReq) error
	Detail(c *gin.Context, req *dtouser.UserDetailReq) (*dtouser.UserDetailResp, error)
	PageList(c *gin.Context, req *dtouser.UserPageListReq) (*dtouser.UserPageListResp, error)
}

type userSvc struct {
}

var _ UserSvc = (*userSvc)(nil)

func NewUserSvc() UserSvc {
	return &userSvc{}
}

// Create 创建用户
func (svc *userSvc) Create(c *gin.Context, req *dtouser.UserCreateReq) (*dtouser.UserCreateResp, error) {
	userId := gincontext.GetUserID(c)
	insertEntity := &model.UserEntity{
		CompanyID:    req.CompanyID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
		CreatedBy:    userId,
		UpdatedBy:    userId,
	}
	if err := daouser.NewUserDao().Insert(c, insertEntity); err != nil {
		glog.Errorf(c, "[svcuser.UserCreate] daoUser Create fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, code.GetError(code.UserCreateErr)
	}
	return &dtouser.UserCreateResp{
		ID: insertEntity.ID,
	}, nil
}

// Delete 删除用户
func (svc *userSvc) Delete(c *gin.Context, req *dtouser.UserDeleteReq) error {
	userID := gincontext.GetUserID(c)

	if err := daouser.NewUserDao().Delete(c, req.ID, userID); err != nil {
		glog.Errorf(c, "[svcuser.Delete] daoUser Delete fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return code.GetError(code.UserDeleteErr)
	}
	return nil
}

// Update 更新用户
func (svc *userSvc) Update(c *gin.Context, req *dtouser.UserUpdateReq) error {
	userId := gincontext.GetUserID(c)
	updateEntity := &model.UserEntity{
		CompanyID:    req.CompanyID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
		UpdatedBy:    userId,
	}
	if err := daouser.NewUserDao().UpdateById(c, req.ID, updateEntity); err != nil {
		glog.Errorf(c, "[svcuser.Update] daoUser UpdateById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return code.GetError(code.UserUpdateErr)
	}
	return nil
}

// Detail 根据id获取用户
func (svc *userSvc) Detail(c *gin.Context, req *dtouser.UserDetailReq) (*dtouser.UserDetailResp, error) {
	detailEntity, err := daouser.NewUserDao().GetById(c, req.ID)
	if err != nil {
		glog.Errorf(c, "[svcuser.UserDetail] daoUser GetById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, code.GetError(code.UserGetDetailErr)
	}
	// 判断是否存在
	if detailEntity == nil || detailEntity.ID == 0 {
		return nil, code.GetError(code.UserNotExistErr)
	}
	Resp := &dtouser.UserDetailResp{
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
	return Resp, nil
}

// PageList 分页获取用户列表
func (svc *userSvc) PageList(c *gin.Context, req *dtouser.UserPageListReq) (*dtouser.UserPageListResp, error) {
	cond := &daouser.UserCond{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	dataList, total, err := daouser.NewUserDao().GetPageListByCond(c, cond)
	if err != nil {
		glog.Errorf(c, "[svcuser.UserPageList] daoUser GetPageListByCond fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, code.GetError(code.UserGetPageListErr)
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
