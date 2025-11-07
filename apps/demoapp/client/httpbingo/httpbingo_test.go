package httpbingo

import (
	"testing"

	"github.com/morehao/go-gin-web/pkg/test"
	"github.com/morehao/golib/glog"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	test.Init("demoapp")

	ctx := test.NewCtx()
	res, err := Get(ctx, &GetRequest{
		ID: 1,
	})
	assert.Nil(t, err)
	t.Log(glog.ToJsonString(res))
}
