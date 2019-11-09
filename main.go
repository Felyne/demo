package main

import (
	"github.com/Felyne/demo/controller"
	_ "github.com/Felyne/demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server demo server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1

func main() {
	port := ":8080"
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":name", c.GetAccount)
			accounts.POST("", c.AddAccount)
			accounts.GET("", c.ListAccounts)
			accounts.DELETE(":name", c.DeleteAccount)
		}
	}
	// 先用swag init 生成docs目录
	// http://localhost:8080/swagger/index.html 也可以swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(port)
}
