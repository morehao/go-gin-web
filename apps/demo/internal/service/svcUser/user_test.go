package svcUser

import (
	"testing"

	"go-gin-web/apps/demo/internal/dto/dtoUser"
	"go-gin-web/apps/demo/internal/object/objUser"
	"go-gin-web/pkg/test"

	"github.com/morehao/go-tools/gutils"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	test.Init(test.AppNameDemo)
	defer test.Done()
	ctx := test.NewCtx(test.OptUid(1))
	req := &dtoUser.UserCreateReq{
		UserBaseInfo: objUser.UserBaseInfo{
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
	req := &dtoUser.UserDetailReq{
		ID: 1,
	}
	resp, err := NewUserSvc().Detail(ctx, req)
	assert.Nil(t, err)
	t.Log(gutils.ToJsonString(resp))
}
