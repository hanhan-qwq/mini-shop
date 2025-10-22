package router

import (
	"github.com/gin-gonic/gin"
	"mini_shop/middleware"
	"mini_shop/web/controller"
)

func RegisterAdminProductRoutes(admin *gin.RouterGroup) {
	productCtrl := controller.NewProductController()

	adminProductGroup := admin.Group("/product")
	{
		admin.Use(middleware.JWTAuthMiddleware("admin"))
		{
			adminProductGroup.POST("", productCtrl.CreateProduct)       // 添加商品
			adminProductGroup.PUT("/:id", productCtrl.UpdateProduct)    // 修改商品
			adminProductGroup.DELETE("/:id", productCtrl.DeleteProduct) // 删除商品
		}
	}
}
