package svc{{.PackagePascalName}}

import (
    "{{.ImportDirPrefix}}/component/errorCode"
	"github.com/morehao/go-tools/gutils"
	"github.com/morehao/go-tools/glog"
	"{{.ImportDirPrefix}}/dto/dto{{.PackagePascalName}}"
	"{{.ImportDirPrefix}}/model/dao{{.PackagePascalName}}"

    "github.com/gin-gonic/gin"
)

// Create 创建{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Create(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}CreateReq) (*dto{{.PackagePascalName}}.{{.StructName}}CreateRes, error) {
	insertEntity := &dao{{.PackagePascalName}}.{{.StructName}}{}
	if err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().Insert(c, insertEntity); err != nil {
		glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}Create] dao{{.PackagePascalName}} Create fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.{{.StructName}}CreateErr
	}
	return &dto{{.PackagePascalName}}.{{.StructName}}CreateRes{
		Id: insertEntity.Id,
	}, nil
}

// Update 更新{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Update(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}UpdateReq) error {
	updateEntity := &dao{{.PackagePascalName}}.{{.StructName}}{
        Id:   req.Id,
    }
    if err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().Update(c, updateEntity); err != nil {
        glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}Update] dao{{.PackagePascalName}} Update fail, err:%v, req:%s", err, gutils.ToJsonString(req))
        return errorCode.{{.StructName}}UpdateErr
    }
    return nil
}

// Detail 根据id获取{{.Description}}
func (svc *{{.ReceiverTypeName}}Svc) Detail(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}DetailReq) (*dto{{.PackagePascalName}}.{{.StructName}}DetailRes, error) {
	detailEntity, err := dao{{.PackagePascalName}}.New{{.StructName}}Dao().GetById(c, req.Id)
	if err != nil {
		glog.Errorf(c, "[svc{{.PackagePascalName}}.{{.StructName}}Detail] dao{{.PackagePascalName}} GetById fail, err:%v, req:%s", err, gutils.ToJsonString(req))
		return nil, errorCode.{{.StructName}}GetDetailErr
	}
    // 判断是否存在
    if detailEntity == nil || detailEntity.Id == 0 {
        return nil, errorCode.{{.StructName}}NotExistErr
    }
	res := &dto{{.PackagePascalName}}.{{.StructName}}DetailRes{
		Id:   detailEntity.Id,
	}
	return res, nil
}

// PageList 分页获取{{.Description}}列表
func (svc *{{.ReceiverTypeName}}Svc) PageList(c *gin.Context, req *dto{{.PackagePascalName}}.{{.StructName}}PageListReq) (*dto{{.PackagePascalName}}.{{.StructName}}PageListRes, error) {
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
	return &dto{{.PackagePascalName}}.{{.StructName}}PageListRes{
		List:  list,
		Total: total,
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
