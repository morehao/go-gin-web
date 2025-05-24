package daouser

import (
	"testing"

	"go-gin-web/internal/apps/demoapp/model"
	"go-gin-web/pkg/test"

	"github.com/morehao/golib/gutils"
	"github.com/stretchr/testify/assert"
)

func TestUserInsert(t *testing.T) {
	test.Init(test.AppNameDemo)
	defer test.Done()
	ctx := test.NewCtx(test.OptUid(1))
	userEntity := &model.UserEntity{
		CompanyID:    1,
		DepartmentID: 1,
		Name:         "test",
	}
	err := NewUserDao().Insert(ctx, userEntity)
	assert.Nil(t, err)
	t.Log(gutils.ToJsonString(userEntity))
}
