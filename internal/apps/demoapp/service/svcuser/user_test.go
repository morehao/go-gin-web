package svcuser

import (
	"testing"

	"go-gin-web/internal/apps/demoapp/dto/dtouser"
	"go-gin-web/internal/apps/demoapp/object/objuser"
	"go-gin-web/pkg/test"

	"github.com/morehao/golib/gutils"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	test.Init(test.AppNameDemo)
	defer test.Done()
	ctx := test.NewCtx(test.OptUid(1))
	req := &dtouser.UserCreateReq{
		UserBaseInfo: objuser.UserBaseInfo{
			CompanyID: 1,
		},
	}
	resp, err := NewUserSvc().Create(ctx, req)
	assert.Nil(t, err)
	t.Log(gutils.ToJsonString(resp))
}

func TestUserDetail(t *testing.T) {
	test.Init(test.AppNameDemo)
	defer test.Done()
	ctx := test.NewCtx(test.OptUid(1))
	req := &dtouser.UserDetailReq{
		ID: 1,
	}
	resp, err := NewUserSvc().Detail(ctx, req)
	assert.Nil(t, err)
	t.Log(gutils.ToJsonString(resp))
}
