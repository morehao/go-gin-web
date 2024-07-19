package svc{{.PackagePascalName}}

import (
	"{{.ImportDirPrefix}}/dto/dto{{.PackagePascalName}}"
	"{{.ImportDirPrefix}}/model/dao{{.PackagePascalName}}"
	"{{.ImportDirPrefix}}/pkg/errorCode"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
	"github.com/morehao/go-tools/gutils"
)

type {{.ReceiverTypePascalName}}Svc interface {
	Create(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}CreateReq) (*dto{{.PackagePascalName}}.{{.StructName}}CreateResp, error)
	Delete(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}DeleteReq) error
	Update(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}UpdateReq) error
	Detail(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}DetailReq) (*dto{{.PackagePascalName}}.{{.StructName}}DetailResp, error)
	PageList(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}PageListReq) (*dto{{.PackagePascalName}}.{{.StructName}}PageListResp, error)
}

type {{.ReceiverTypeName}}Svc struct {
}

var _ {{.ReceiverTypePascalName}}Svc = (*{{.ReceiverTypeName}}Svc)(nil)

func New{{.ReceiverTypePascalName}}Svc() {{.ReceiverTypePascalName}}Svc {
	return &{{.ReceiverTypeName}}Svc{}
}

// Create 创建{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Create(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}CreateReq) (*dto{{.PackagePascalName}}.{{.StructName}}CreateResp, error) {
	insertEntity := &dao{{.PackagePascalName}}.{{.StructName}}Entity{}
	if err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().Insert(c, insertEntity); err != nil {
		glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}Create] dao{{.PackagePascalName}} Create fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.{{.StructName}}CreateErr
	}
	return &dto{{.PackagePascalName}}.{{.StructName}}CreateResp{
		Id: insertEntity.Id,
	}, nil
}

// Delete 删除{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Delete(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}DeleteReq) error {
	deletedBy := context.GetUserId(c)

	if err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().Delete(c, req.Id, deletedBy); err != nil {
		glog.Errorf(c, "[svc{{.PackagePascalName}}.Delete] dao{{.PackagePascalName}} Delete fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return errorCode.{{.StructName}}DeleteErr
	}
	return nil
}

// Update 更新{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Update(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}UpdateReq) error {
	updateEntity := &dao{{.PackagePascalName}}.{{.StructName}}Entity{
        Id:   req.Id,
    }
    if err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().Update(c, updateEntity); err != nil {
        glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}Update] dao{{.PackagePascalName}} Update fail, err:%v, req:%s", err, gutils.ToJsonString(req))
        return errorCode.{{.StructName}}UpdateErr
    }
    return nil
}

// Detail 根据id获取{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Detail(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}DetailReq) (*dto{{.PackagePascalName}}.{{.StructName}}DetailResp, error) {
	detailEntity, err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().GetById(c, req.Id)
	if err != nil {
		glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}Detail] dao{{.PackagePascalName}} GetById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.{{.StructName}}GetDetailErr
	}
    // 判断是否存在
    if detailEntity == nil || detailEntity.Id == 0 {
        return nil, errorCode.{{.StructName}}NotExistErr
    }
	Resp := &dto{{.PackagePascalName}}.{{.StructName}}DetailResp{
		Id:   detailEntity.Id,
	}
	return Resp, nil
}

// PageList 分页获取{{.Description}}列表
func (svc *{{.ReceiverTypeName}}Svc) PageList(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}PageListReq) (*dto{{.PackagePascalName}}.{{.StructName}}PageListResp, error) {
	cond := &dao{{.PackagePascalName}}.{{.StructName}}Cond{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	dataList, total, err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().GetPageListByCond(c, cond)
	if err != nil {
		glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}PageList] dao{{.PackagePascalName}} GetPageListByCond fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.{{.StructName}}GetPageListErr
	}
	list := make([]dto{{.PackagePascalName}}.{{.StructName}}PageListItem, 0, len(dataList))
	for _, v := range dataList {
		list = append(list, dto{{.PackagePascalName}}.{{.StructName}}PageListItem{
			Id:        v.Id,
			CreatedAt: v.CreatedAt.Unix(),
		})
	}
	return &dto{{.PackagePascalName}}.{{.StructName}}PageListResp{
		List:  list,
		Total: total,
	}, nil
}


