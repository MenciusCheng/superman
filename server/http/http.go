package http

import (
	"github.com/MenciusCheng/superman/conf"
	"github.com/MenciusCheng/superman/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	svc *service.Service

	httpServer *gin.Engine
)

// Init create a rpc server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	// new http server
	httpServer = gin.New()

	// add namespace plugin
	httpServer.Use(Logger())

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("http server start failed, err %v", err)
		}
	}()

}

func Shutdown() {
	//if httpServer != nil {
	//	httpServer.Stop()
	//}
	if svc != nil {
		svc.Close()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
