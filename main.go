package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tlpazmt/goProject/api/controllers"
	"github.com/tlpazmt/goProject/api/middlewares"
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
	r.POST("/login", controllers.Login)
	r.GET("/books", middlewares.Auth, controllers.GetBooks)
	r.POST("/books", middlewares.Auth, controllers.CreateBook)
	r.POST("books/:id/rating", middlewares.Auth, controllers.SetBookRating)
	r.POST("books/:id/comment", middlewares.Auth, controllers.CreateComment)
	r.POST("books/:id/purchase", middlewares.Auth, controllers.PurchaseBook)

	r.Run()
}
