package main

import (
	"github.com/DaHuangQwQ/ginx"
	"github.com/DaHuangQwQ/ginx/middleware/prometheus"
	"github.com/gin-gonic/gin"
)

func main() {
	server := ginx.NewServer(":8081")

	builder := prometheus.Builder{
		Namespace:  "test",
		Subsystem:  "test",
		Name:       "user",
		InstanceId: "1",
		Help:       "1",
	}
	server.Use(builder.BuildActiveRequest())
	server.Use(builder.BuildResponseTime())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, ginx.Result[any]{
			Code: 0,
			Msg:  "ok",
			Data: "hello world",
		})
	})
	_ = server.Start()
}
