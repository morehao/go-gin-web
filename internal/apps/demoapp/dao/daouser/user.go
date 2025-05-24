package daouser

import (
	"fmt"
	"time"

	"go-gin-web/internal/apps/demoapp/code"
	"go-gin-web/internal/apps/demoapp/dao"
	"go-gin-web/internal/apps/demoapp/model"

	"github.com/gin-gonic/gin"
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

func (repo *UserDao) TableName() string {
	return model.TblNameUser
}

func (repo *UserDao) WithTx(db *gorm.DB) *UserDao {
	return &UserDao{
		Base: dao.Base{Tx: db},
	}
}

func (repo *UserDao) Insert(c *gin.Context, entity *model.UserEntity) error {
	db := repo.Db(c).Omit().Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.ErrorDbInsert.Wrapf(err, "[UserDao] Insert fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (repo *UserDao) BatchInsert(c *gin.Context, entityList model.UserEntityList) error {
	db := repo.Db(c).Table(repo.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.ErrorDbInsert.Wrapf(err, "[UserDao] BatchInsert fail, entityList:%s", gutils.ToJsonString(entityList))
	}
	return nil
}

func (repo *UserDao) UpdateById(c *gin.Context, id uint, entity *model.UserEntity) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.ErrorDbUpdate.Wrapf(err, "[UserDao] UpdateById fail, id:%d, entity:%s", id, gutils.ToJsonString(entity))
	}
	return nil
}

func (repo *UserDao) UpdateMap(c *gin.Context, id uint, updateMap map[string]interface{}) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.ErrorDbUpdate.Wrapf(err, "[UserDao] UpdateMap fail, id:%d, updateMap:%s", id, gutils.ToJsonString(updateMap))
	}
	return nil
}

func (repo *UserDao) Delete(c *gin.Context, id, deletedBy uint) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.ErrorDbUpdate.Wrapf(err, "[UserDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (repo *UserDao) GetById(c *gin.Context, id uint) (*model.UserEntity, error) {
	var entity model.UserEntity
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.ErrorDbFind.Wrapf(err, "[UserDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (repo *UserDao) GetByCond(c *gin.Context, cond *UserCond) (*model.UserEntity, error) {
	var entity model.UserEntity
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())

	repo.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.ErrorDbFind.Wrapf(err, "[UserDao] GetById fail, cond:%s", gutils.ToJsonString(cond))
	}
	return &entity, nil
}

func (repo *UserDao) GetListByCond(c *gin.Context, cond *UserCond) (model.UserEntityList, error) {
	var entityList model.UserEntityList
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())

	repo.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.ErrorDbFind.Wrapf(err, "[UserDao] GetListByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, nil
}

func (repo *UserDao) GetPageListByCond(c *gin.Context, cond *UserCond) (model.UserEntityList, int64, error) {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())

	repo.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.ErrorDbFind.Wrapf(err, "[UserDao] GetPageListByCond count fail, cond:%s", gutils.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var list model.UserEntityList
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, code.ErrorDbFind.Wrapf(err, "[UserDao] GetPageListByCond find fail, cond:%s", gutils.ToJsonString(cond))
	}
	return list, count, nil
}

func (repo *UserDao) BuildCondition(db *gorm.DB, cond *UserCond) {
	if cond.ID > 0 {
		query := fmt.Sprintf("%s.id = ?", repo.TableName())
		db.Where(query, cond.ID)
	}
	if len(cond.IDs) > 0 {
		query := fmt.Sprintf("%s.id in (?)", repo.TableName())
		db.Where(query, cond.IDs)
	}
	if cond.CreatedAtStart > 0 {
		query := fmt.Sprintf("%s.created_at >= ?", repo.TableName())
		db.Where(query, time.Unix(cond.CreatedAtStart, 0))
	}
	if cond.CreatedAtEnd > 0 {
		query := fmt.Sprintf("%s.created_at <= ?", repo.TableName())
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
