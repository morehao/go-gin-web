package svcUser

import (
	"strings"

	"go-gin-web/apps/demo/internal/dto/dtoUser"
	"go-gin-web/apps/demo/internal/errorCode"
	"go-gin-web/apps/demo/internal/model/mysqlModel/daoUser"
	"go-gin-web/apps/demo/internal/object/objCommon"
	"go-gin-web/apps/demo/internal/object/objUser"
	"go-gin-web/pkg/cuctx"
	"go-gin-web/pkg/storages"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
	"github.com/morehao/go-tools/gutils"
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
	userId := cuctx.GetUserID(c)
	now := time.Now()
	insertEntity := &daoUser.UserEntity{
		CompanyID:    req.CompanyID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
		CreatedBy:    userId,
		CreatedAt:    now,
		UpdatedBy:    userId,
		UpdatedAt:    now,
	}
	if err := daoUser.NewUserDao().Insert(c, insertEntity); err != nil {
		glog.Errorf(c, "[svcUser.UserCreate] daoUser Create fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.UserCreateErr
	}
	return &dtoUser.UserCreateResp{
		ID: insertEntity.ID,
	}, nil
}

// Delete 删除用户
func (svc *userSvc) Delete(c *gin.Context, req *dtoUser.UserDeleteReq) error {
	deletedBy := cuctx.GetUserID(c)

	if err := daoUser.NewUserDao().Delete(c, req.ID, deletedBy); err != nil {
		glog.Errorf(c, "[svcUser.Delete] daoUser Delete fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return errorCode.UserDeleteErr
	}
	return nil
}

// Update 更新用户
func (svc *userSvc) Update(c *gin.Context, req *dtoUser.UserUpdateReq) error {
	updateEntity := &daoUser.UserEntity{
		ID: req.ID,
	}
	if err := daoUser.NewUserDao().Update(c, updateEntity); err != nil {
		glog.Errorf(c, "[svcUser.UserUpdate] daoUser Update fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return errorCode.UserUpdateErr
	}
	return nil
}

// Detail 根据id获取用户
func (svc *userSvc) Detail(c *gin.Context, req *dtoUser.UserDetailReq) (*dtoUser.UserDetailResp, error) {
	detailEntity, err := daoUser.NewUserDao().GetById(c, req.ID)
	if err != nil {
		glog.Errorf(c, "[svcUser.UserDetail] daoUser GetById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.UserGetDetailErr
	}
	// 判断是否存在
	if detailEntity == nil || detailEntity.ID == 0 {
		return nil, errorCode.UserNotExistErr
	}
	Resp := &dtoUser.UserDetailResp{
		ID: detailEntity.ID,
		UserBaseInfo: objUser.UserBaseInfo{
			CompanyID:    detailEntity.CompanyID,
			DepartmentID: detailEntity.DepartmentID,
			Name:         detailEntity.Name,
		},
		OperatorBaseInfo: objCommon.OperatorBaseInfo{
			CreatedBy: detailEntity.CreatedBy,
			CreatedAt: detailEntity.CreatedAt.Unix(),
			UpdatedBy: detailEntity.UpdatedBy,
			UpdatedAt: detailEntity.UpdatedAt.Unix(),
		},
	}
	return Resp, nil
}

// PageList 分页获取用户列表
func (svc *userSvc) PageList(c *gin.Context, req *dtoUser.UserPageListReq) (*dtoUser.UserPageListResp, error) {
	// 为了测试各组件的日志
	glog.Infof(c, "[svcUser.UserPageList] req:%s", glog.ToJsonString(req))
	_, _ = storages.DemoRedis.Get(c, "test").Result()
	_, _ = storages.DemoES.Search(
		storages.DemoES.Search.WithContext(c),
		storages.DemoES.Search.WithIndex("accounts"),
		storages.DemoES.Search.WithBody(strings.NewReader(`{"query":{"match_all":{}}}`)),
	)

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
			ID: v.ID,
			UserBaseInfo: objUser.UserBaseInfo{
				CompanyID:    v.CompanyID,
				DepartmentID: v.DepartmentID,
				Name:         v.Name,
			},
			OperatorBaseInfo: objCommon.OperatorBaseInfo{
				CreatedBy: v.CreatedBy,
				CreatedAt: v.CreatedAt.Unix(),
				UpdatedBy: v.UpdatedBy,
				UpdatedAt: v.UpdatedAt.Unix(),
			},
		})
	}
	return &dtoUser.UserPageListResp{
		List:  list,
		Total: total,
	}, nil
}
