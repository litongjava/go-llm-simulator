package controller

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func PingHandler(c context.Context, ctx *app.RequestContext) {
  ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
}
