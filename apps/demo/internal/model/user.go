package model

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID           uint64         `gorm:"column:id;comment:自增ID;primaryKey"`
	CompanyID    uint64         `gorm:"column:company_id;comment:公司id"`
	DepartmentID uint64         `gorm:"column:department_id;comment:部门id"`
	Name         string         `gorm:"column:name;comment:姓名"`
	CreatedAt    time.Time      `gorm:"column:created_at;comment:创建时间"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;comment:更新时间"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	CreatedBy    uint64         `gorm:"column:created_by;comment:创建人id"`
	UpdatedBy    uint64         `gorm:"column:updated_by;comment:更新人id"`
	DeletedBy    uint64         `gorm:"column:deleted_by;comment:删除人id"`
}

type UserEntityList []UserEntity

const TblNameUser = "user"

func (UserEntity) TableName() string {
	return TblNameUser
}

func (l UserEntityList) ToMap() map[uint64]UserEntity {
	m := make(map[uint64]UserEntity)
	for _, v := range l {
		m[v.ID] = v
	}
	return m
}
