package httpbingo

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-gin-web/apps/demoapp/config"
	"github.com/morehao/go-gin-web/apps/demoapp/dao/daouser"
	"github.com/morehao/go-gin-web/pkg/storages"
)

func Get(ctx *gin.Context, req *GetRequest) (*GetResponse, error) {
	_, _ = storages.DemoRedis.Get(ctx, "").Result()
	_, _ = daouser.NewUserDao().GetListByCond(ctx, &daouser.UserCond{})
	var res GetResponse
	cfg := config.Conf
	request, getRequestErr := cfg.Client.HTTPBingo.NewRequestWithResult(ctx, &res)
	if getRequestErr != nil {
		return nil, getRequestErr
	}
	params := map[string]string{
		"id": strconv.FormatUint(uint64(req.ID), 10),
	}
	aa, getErr := request.SetQueryParams(params).Get("/get")
	if getErr != nil {
		return nil, getErr
	}
	fmt.Println(aa)
	return &res, nil
}
