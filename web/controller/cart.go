package controller

import (
	"fmt"
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

// CreateItem POST /api/v1/cart
func (ctrl *CartController) CreateItem(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "用户未登录",
		})
		return
	}

	var req request.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
		return
	}

	fmt.Print(userID, req.ProductID)
	err := ctrl.CartService.CreateItem(userID.(uint), req.ProductID, req.Quantity)
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

// GetCart Get /api/v1/cart
func (ctrl *CartController) GetCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "用户未登录",
		})
	}

	items, total, err := ctrl.CartService.GetCart(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "获取购物车列表失败",
			"error":   err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        "0",
		"message":     "成功获取购物车信息",
		"data":        items,
		"total_price": total,
	})
}

// UpdateItem PUT /api/v1/cart/update
func (ctrl *CartController) UpdateItem(c *gin.Context) {
	var req request.UpdateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "用户未登录",
		})
	}

	err := ctrl.CartService.UpdateItem(userID.(uint), req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "更新购物车失败",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "更新购物车成功",
	})
}

func (ctrl *CartController) DeleteItem(c *gin.Context) {
	var req request.DeleteCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "用户未登录",
		})
	}

	err := ctrl.CartService.DeleteItem(userID.(uint), req.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "商品从购物车移除失败",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "商品成功从购物车移除",
	})
}
