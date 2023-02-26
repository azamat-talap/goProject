package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tlpazmt/goProject/api/controllers"
	"github.com/tlpazmt/goProject/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)

	r.Run()
}