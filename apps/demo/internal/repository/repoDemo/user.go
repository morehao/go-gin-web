package repoDemo

import (
	"fmt"
	"time"

	"go-gin-web/apps/demo/internal/errorCode"
	"go-gin-web/apps/demo/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gutils"
	"gorm.io/gorm"
)

type UserCond struct {
	ID             uint64
	IDs            []uint64
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type UserRepo struct {
	Base
}

func NewUser() *UserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) TableName() string {
	return model.TblNameUser
}

func (repo *UserRepo) WithTx(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Base: Base{Tx: db},
	}
}

func (repo *UserRepo) Insert(c *gin.Context, entity *model.UserEntity) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Create(entity).Error; err != nil {
		return errorCode.ErrorDbInsert.Wrapf(err, "[UserRepo] Insert fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (repo *UserRepo) BatchInsert(c *gin.Context, entityList model.UserEntityList) error {
	db := repo.Db(c).Table(repo.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return errorCode.ErrorDbInsert.Wrapf(err, "[UserRepo] BatchInsert fail, entityList:%s", gutils.ToJsonString(entityList))
	}
	return nil
}

func (repo *UserRepo) Update(c *gin.Context, entity *model.UserEntity) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Where("id = ?", entity.ID).Updates(entity).Error; err != nil {
		return errorCode.ErrorDbUpdate.Wrapf(err, "[UserRepo] Update fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (repo *UserRepo) UpdateMap(c *gin.Context, id uint64, updateMap map[string]interface{}) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return errorCode.ErrorDbUpdate.Wrapf(err, "[UserRepo] UpdateMap fail, id:%d, updateMap:%s", id, gutils.ToJsonString(updateMap))
	}
	return nil
}

func (repo *UserRepo) Delete(c *gin.Context, id, deletedBy uint64) error {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return errorCode.ErrorDbUpdate.Wrapf(err, "[UserRepo] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (repo *UserRepo) GetById(c *gin.Context, id uint64) (*model.UserEntity, error) {
	var entity model.UserEntity
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, errorCode.ErrorDbFind.Wrapf(err, "[UserRepo] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (repo *UserRepo) GetByCond(c *gin.Context, cond *UserCond) (*model.UserEntity, error) {
	var entity model.UserEntity
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())

	repo.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, errorCode.ErrorDbFind.Wrapf(err, "[UserRepo] GetById fail, cond:%s", gutils.ToJsonString(cond))
	}
	return &entity, nil
}

func (repo *UserRepo) GetListByCond(c *gin.Context, cond *UserCond) (model.UserEntityList, error) {
	var entityList model.UserEntityList
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())

	repo.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, errorCode.ErrorDbFind.Wrapf(err, "[UserRepo] GetListByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, nil
}

func (repo *UserRepo) GetPageListByCond(c *gin.Context, cond *UserCond) (model.UserEntityList, int64, error) {
	db := repo.Db(c).Model(&model.UserEntity{})
	db = db.Table(repo.TableName())

	repo.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, errorCode.ErrorDbFind.Wrapf(err, "[UserRepo] GetPageListByCond count fail, cond:%s", gutils.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var list model.UserEntityList
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, errorCode.ErrorDbFind.Wrapf(err, "[UserRepo] GetPageListByCond find fail, cond:%s", gutils.ToJsonString(cond))
	}
	return list, count, nil
}

func (repo *UserRepo) BuildCondition(db *gorm.DB, cond *UserCond) {
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
