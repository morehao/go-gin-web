package daouser

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-gin-web/apps/demoapp/dao"
	"github.com/morehao/go-gin-web/apps/demoapp/model"
	"github.com/morehao/go-gin-web/pkg/code"
	"github.com/morehao/golib/gutils"
	"gorm.io/gorm"
)

type UserCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type UserDao struct {
	dao.Base
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) TableName() string {
	return model.TableNameUser
}

func (d *UserDao) WithTx(db *gorm.DB) *UserDao {
	return &UserDao{
		Base: dao.Base{Tx: db},
	}
}

func (d *UserDao) Insert(ctx *gin.Context, entity *model.UserEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(code.DBInsertErr).Wrapf(err, "[UserDao] Insert fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (d *UserDao) BatchInsert(ctx *gin.Context, entityList model.UserEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(code.DBInsertErr).Wrapf(nil, "[UserDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(code.DBInsertErr).Wrapf(err, "[UserDao] BatchInsert fail, entityList:%s", gutils.ToJsonString(entityList))
	}
	return nil
}

func (d *UserDao) UpdateByID(ctx *gin.Context, id uint, entity *model.UserEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[UserDao] UpdateByID fail, id:%d entity:%s", id, gutils.ToJsonString(entity))
	}
	return nil
}

func (d *UserDao) UpdateMap(ctx *gin.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[UserDao] UpdateMap fail, id:%d, updateMap:%s", id, gutils.ToJsonString(updateMap))
	}
	return nil
}

func (d *UserDao) Delete(ctx *gin.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[UserDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *UserDao) GetById(ctx *gin.Context, id uint) (*model.UserEntity, error) {
	var entity model.UserEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[UserDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *UserDao) GetByCond(ctx *gin.Context, cond *UserCond) (*model.UserEntity, error) {
	var entity model.UserEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[UserDao] GetById fail, cond:%s", gutils.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *UserDao) GetListByCond(ctx *gin.Context, cond *UserCond) (model.UserEntityList, error) {
	var entityList model.UserEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[UserDao] GetListByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *UserDao) GetPageListByCond(ctx *gin.Context, cond *UserCond) (model.UserEntityList, int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(code.DBFindErr).Wrapf(err, "[UserDao] GetPageListByCond count fail, cond:%s", gutils.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList model.UserEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(code.DBFindErr).Wrapf(err, "[UserDao] GetPageListByCond find fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *UserDao) CountByCond(ctx *gin.Context, cond *UserCond) (int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(code.DBFindErr).Wrapf(err, "[UserDao] CountByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return count, nil
}

func (d *UserDao) BuildCondition(db *gorm.DB, cond *UserCond) {
	if cond.ID > 0 {
		query := fmt.Sprintf("%s.id = ?", d.TableName())
		db.Where(query, cond.ID)
	}
	if len(cond.IDs) > 0 {
		query := fmt.Sprintf("%s.id in (?)", d.TableName())
		db.Where(query, cond.IDs)
	}
	if cond.CreatedAtStart > 0 {
		query := fmt.Sprintf("%s.created_at >= ?", d.TableName())
		db.Where(query, time.Unix(cond.CreatedAtStart, 0))
	}
	if cond.CreatedAtEnd > 0 {
		query := fmt.Sprintf("%s.created_at <= ?", d.TableName())
		db.Where(query, time.Unix(cond.CreatedAtEnd, 0))
	}
	if cond.IsDelete {
		db.Unscoped()
	}

	if cond.OrderField != "" {
		db.Order(cond.OrderField)
	}

	return
}
