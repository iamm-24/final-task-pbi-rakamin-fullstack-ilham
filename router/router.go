package router

import (
	"rakamin/controllers"
	middlewares "rakamin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users/register", controllers.RegisterUser)
		api.POST("/users/login", controllers.LoginUser)
		api.PUT("/users/:userId", middlewares.JWTAuth(), controllers.UpdateUser)
		api.DELETE("/users/:userId", middlewares.JWTAuth(), controllers.DeleteUser)

		api.POST("/photos", middlewares.JWTAuth(), controllers.AddPhoto)
		api.GET("/photos", middlewares.JWTAuth(), controllers.GetPhotos)
		api.PUT("/photos/:photoId", middlewares.JWTAuth(), controllers.UpdatePhoto)
		api.DELETE("/photos/:photoId", middlewares.JWTAuth(), controllers.DeletePhoto)
	}

	return r
}
