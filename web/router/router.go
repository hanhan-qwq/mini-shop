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

		productCtrl := controller.NewProductController()
		productGroup := v1.Group("/product")
		{
			productGroup.GET("", productCtrl.ListProducts)
			productGroup.GET("/:id", productCtrl.GetProduct)
		}

		admin := v1.Group("/admin")
		{
			RegisterAdminProductRoutes(admin)
			//RegisterAdminOrderRoutes(admin)
		}

		cartCtrl := controller.NewCartController()
		cartGroup := v1.Group("/cart")
		{
			cartGroup.Use(middleware.JWTAuthMiddleware("user", "admin"))
			{
				cartGroup.POST("", cartCtrl.CreateItem)
				cartGroup.GET("", cartCtrl.GetCart)
				cartGroup.PUT("", cartCtrl.UpdateItem)
				cartGroup.DELETE("", cartCtrl.DeleteItem)
			}
		}

		orderCtrl := controller.NewOrderController()
		orderGroup := v1.Group("/order")
		{
			orderGroup.Use(middleware.JWTAuthMiddleware("user", "admin"))
			{
				orderGroup.POST("/create", orderCtrl.CreateOrder)
			}
		}
	}

	return r
}
