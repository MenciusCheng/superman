package http

import (
	"context"
	"github.com/MenciusCheng/superman/util/ecode"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	if err := svc.Ping(context.Background()); err != nil {
		c.JSON(ecode.Cause(err).Code(), err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(200, okMsg)
}
