package svcUser

import (
	"testing"

	"go-gin-web/apps/demo/dto/dtoUser"
	"go-gin-web/apps/demo/object/objUser"
	test2 "go-gin-web/pkg/test"

	"github.com/morehao/go-tools/gutils"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	test2.Init()
	defer test2.Done()
	ctx := test2.NewCtx(test2.OptUid(1))
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
	test2.Init()
	defer test2.Done()
	ctx := test2.NewCtx(test2.OptUid(1))
	req := &dtoUser.UserDetailReq{
		ID: 1,
	}
	resp, err := NewUserSvc().Detail(ctx, req)
	assert.Nil(t, err)
	t.Log(gutils.ToJsonString(resp))
}
