package router

import (
	"github.com/gin-gonic/gin"
	"mini_shop/middleware"
	"mini_shop/web/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controller.UserRegister)
			auth.POST("/login", controller.UserLogin)
			//auth.POST("/logout",controller.UserLogout)
		}
		user := v1.Group("/user")
		{
			user.Use(middleware.JWTAuthMiddleware("user", "admin"))
			{
				user.GET("/profile", controller.GetUserProfile)
				user.PUT("/profile", controller.UpdateUserProfile)
				//user.DELETE("/profile",controller.DeleteUserProfile)
			}
		}
	}

	return r
}
