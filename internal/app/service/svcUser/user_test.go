package svcUser

import (
	"testing"

	"go-gin-web/internal/app/dto/dtoUser"
	"go-gin-web/internal/pkg/test"

	"github.com/morehao/go-tools/gutils"
	"github.com/stretchr/testify/assert"
)

func TestUserDetail(t *testing.T) {
	test.Init()
	defer test.Done()
	ctx := test.NewCtx(test.OptUid(1))
	req := &dtoUser.UserDetailReq{
		ID: 1,
	}
	resp, err := NewUserSvc().Detail(ctx, req)
	assert.Nil(t, err)
	t.Log(gutils.ToJsonString(resp))
}
