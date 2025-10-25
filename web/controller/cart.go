package controller

import (
	"github.com/gin-gonic/gin"
	"mini_shop/service"
	"mini_shop/web/request"
	"net/http"
)

type CartController struct {
	CartService *service.CartService
}

func NewCartController() *CartController {
	return &CartController{
		CartService: service.NewCartService(),
	}
}

// AddToCart POST /api/v1/cart
func (ctrl *CartController) AddToCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "用户未登录",
		})
		return
	}

	var req request.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
		return
	}

	err := ctrl.CartService.AddToCart(userId.(uint), req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "添加购物车失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "添加购物车成功",
	})
}
