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
		authCtrl := controller.NewAuthController()
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", authCtrl.UserRegister)
			authGroup.POST("/login", authCtrl.UserLogin)
			//auth.POST("/logout",controller.UserLogout)
		}

		userCtrl := controller.NewUserController()
		userGroup := v1.Group("/user")
		{
			userGroup.Use(middleware.JWTAuthMiddleware("user", "admin"))
			{
				userGroup.GET("/profile", userCtrl.GetUserProfile)
				userGroup.PUT("/profile", userCtrl.UpdateUserProfile)
				//user.DELETE("/profile",controller.DeleteUserProfile)
				//user.GET("/orders", controller.ListOrders)
			}
		}
	}

	return r
}
