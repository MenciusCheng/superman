package http

import (
	"github.com/gin-gonic/gin"
)

func initRoute(s *gin.Engine) {

	s.GET("/ping", ping)

}
