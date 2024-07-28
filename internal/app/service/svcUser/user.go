package svcUser

import (
	"go-gin-web/internal/app/dto/dtoUser"
	"go-gin-web/internal/app/model/daoUser"
	"go-gin-web/internal/app/object/objCommon"
	"go-gin-web/internal/app/object/objUser"
	"go-gin-web/internal/pkg/context"
	"go-gin-web/internal/pkg/errorCode"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
	"github.com/morehao/go-tools/gutils"
	"time"
)

type UserSvc interface {
	Create(c *gin.Context, req *dtoUser.UserCreateReq) (*dtoUser.UserCreateResp, error)
	Delete(c *gin.Context, req *dtoUser.UserDeleteReq) error
	Update(c *gin.Context, req *dtoUser.UserUpdateReq) error
	Detail(c *gin.Context, req *dtoUser.UserDetailReq) (*dtoUser.UserDetailResp, error)
	PageList(c *gin.Context, req *dtoUser.UserPageListReq) (*dtoUser.UserPageListResp, error)
}

type userSvc struct {
}

var _ UserSvc = (*userSvc)(nil)

func NewUserSvc() UserSvc {
	return &userSvc{}
}

// Create 创建用户
func (svc *userSvc) Create(c *gin.Context, req *dtoUser.UserCreateReq) (*dtoUser.UserCreateResp, error) {
	userId := context.GetUserId(c)
	now := time.Now()
	insertEntity := &daoUser.UserEntity{
		CompanyId: req.CompanyId,
		DepartmentId: req.DepartmentId,
		Name: req.Name,
		CreatedBy: userId,
		CreatedTime: now,
		UpdatedBy: userId,
		UpdatedTime: now,
	}
	if err := daoUser.NewUserDao().Insert(c, insertEntity); err != nil {
		glog.Errorf(c, "[svcUser.UserCreate] daoUser Create fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.UserCreateErr
	}
	return &dtoUser.UserCreateResp{
		Id: insertEntity.Id,
	}, nil
}

// Delete 删除用户
func (svc *userSvc) Delete(c *gin.Context, req *dtoUser.UserDeleteReq) error {
	deletedBy := context.GetUserId(c)

	if err := daoUser.NewUserDao().Delete(c, req.Id, deletedBy); err != nil {
		glog.Errorf(c, "[svcUser.Delete] daoUser Delete fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return errorCode.UserDeleteErr
	}
	return nil
}

// Update 更新用户
func (svc *userSvc) Update(c *gin.Context, req *dtoUser.UserUpdateReq) error {
	updateEntity := &daoUser.UserEntity{
        Id:   req.Id,
    }
    if err := daoUser.NewUserDao().Update(c, updateEntity); err != nil {
        glog.Errorf(c, "[svcUser.UserUpdate] daoUser Update fail, err:%v, req:%s", err, gutils.ToJsonString(req))
        return errorCode.UserUpdateErr
    }
    return nil
}

// Detail 根据id获取用户
func (svc *userSvc) Detail(c *gin.Context, req *dtoUser.UserDetailReq) (*dtoUser.UserDetailResp, error) {
	detailEntity, err := daoUser.NewUserDao().GetById(c, req.Id)
	if err != nil {
		glog.Errorf(c, "[svcUser.UserDetail] daoUser GetById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.UserGetDetailErr
	}
    // 判断是否存在
    if detailEntity == nil || detailEntity.Id == 0 {
        return nil, errorCode.UserNotExistErr
    }
	Resp := &dtoUser.UserDetailResp{
		Id:   detailEntity.Id,
		UserBaseInfo: objUser.UserBaseInfo{
			CompanyId: detailEntity.CompanyId,
			DepartmentId: detailEntity.DepartmentId,
			Name: detailEntity.Name,
		},
		OperatorBaseInfo: objCommon.OperatorBaseInfo{
        	CreatedBy:   detailEntity.CreatedBy,
			CreatedTime: detailEntity.CreatedTime.Unix(),
			UpdatedBy:   detailEntity.UpdatedBy,
			UpdatedTime: detailEntity.UpdatedTime.Unix(),
		},
	}
	return Resp, nil
}

// PageList 分页获取用户列表
func (svc *userSvc) PageList(c *gin.Context, req *dtoUser.UserPageListReq) (*dtoUser.UserPageListResp, error) {
	cond := &daoUser.UserCond{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	dataList, total, err := daoUser.NewUserDao().GetPageListByCond(c, cond)
	if err != nil {
		glog.Errorf(c, "[svcUser.UserPageList] daoUser GetPageListByCond fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.UserGetPageListErr
	}
	list := make([]dtoUser.UserPageListItem, 0, len(dataList))
	for _, v := range dataList {
		list = append(list, dtoUser.UserPageListItem{
			Id:   v.Id,
			UserBaseInfo: objUser.UserBaseInfo{
				CompanyId: v.CompanyId,
				DepartmentId: v.DepartmentId,
				Name: v.Name,
			},
			OperatorBaseInfo: objCommon.OperatorBaseInfo{
				CreatedBy:   v.CreatedBy,
				CreatedTime: v.CreatedTime.Unix(),
				UpdatedBy:   v.UpdatedBy,
				UpdatedTime: v.UpdatedTime.Unix(),
			},
		})
	}
	return &dtoUser.UserPageListResp{
		List:  list,
		Total: total,
	}, nil
}


