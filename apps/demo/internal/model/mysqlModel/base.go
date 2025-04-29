package mysqlModel

import (
	"go-gin-web/pkg/storages"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Base struct {
	Tx *gorm.DB
}

// Db 获取Db
func (base *Base) Db(ctx *gin.Context) (db *gorm.DB) {
	if base.Tx != nil {
		return base.Tx.WithContext(ctx)
	}

	db = storages.DBDemo.WithContext(ctx)
	return
}
