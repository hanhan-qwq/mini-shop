package controller

import (
	"github.com/gin-gonic/gin"
	"mini_shop/service"
	"mini_shop/web/request"
	"net/http"
)

// OrderController 订单控制器
type OrderController struct {
	OrderService *service.OrderService // 依赖订单服务层
}

// NewOrderController 创建订单控制器实例
func NewOrderController() *OrderController {
	return &OrderController{
		OrderService: service.NewOrderService(), // 初始化订单服务
	}
}

// CreateOrder 创建订单
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	// 1. 获取当前登录用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "用户未登录",
		})
		return
	}

	// 2. 绑定请求参数
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 3. 调用服务层创建订单
	err := ctrl.OrderService.CreateOrder(
		userID.(uint), // 转换用户ID类型
		req.Remark,
		req.Items,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "创建订单失败",
			"error":   err.Error(),
		})
		return
	}

	// 4. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建订单成功",
	})
}

// GetOrder 查看订单详情
//func (ctrl *OrderController) GetOrder(c *gin.Context) {
//	// 1. 解析订单ID
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil || id <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code":    -1,
//			"message": "订单ID无效",
//		})
//		return
//	}
//
//	// 2. 获取当前用户ID（验证权限）
//	userID, _ := c.Get("user_id")
//
//	// 3. 调用服务层查询订单
//	order, err := ctrl.OrderService.GetOrderDetail(uint(id), userID.(uint))
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{
//			"code":    -1,
//			"message": "查询订单失败",
//			"error":   err.Error(),
//		})
//		return
//	}
//
//	// 4. 返回订单详情
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"data": order,
//	})
//}
//
//// ListOrders 分页查询用户订单列表
//func (ctrl *OrderController) ListOrders(c *gin.Context) {
//	// 1. 获取分页参数
//	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
//	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
//	status, _ := strconv.Atoi(c.DefaultQuery("status", "0")) // 0表示全部状态
//
//	// 2. 获取当前用户ID
//	userID, exists := c.Get("user_id")
//	if !exists {
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"code":    -1,
//			"message": "请先登录",
//		})
//		return
//	}
//
//	// 3. 调用服务层查询订单列表
//	orders, count, err := ctrl.OrderService.ListUserOrders(
//		userID.(uint),
//		page,
//		pageSize,
//		status,
//	)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"code":    -1,
//			"message": "查询订单列表失败",
//			"error":   err.Error(),
//		})
//		return
//	}
//
//	// 4. 返回分页结果
//	c.JSON(http.StatusOK, gin.H{
//		"code":      0,
//		"orders":    orders,
//		"count":     count,
//		"page":      page,
//		"page_size": pageSize,
//	})
//}
//
//// PayOrder 支付订单
//func (ctrl *OrderController) PayOrder(c *gin.Context) {
//	// 1. 解析订单号
//	orderNo := c.Param("order_no")
//	if orderNo == "" {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code":    -1,
//			"message": "订单号不能为空",
//		})
//		return
//	}
//
//	// 2. 获取支付方式（1:微信，2:支付宝）
//	payType, _ := strconv.Atoi(c.DefaultPostForm("pay_type", "1"))
//
//	// 3. 调用服务层处理支付
//	err := ctrl.OrderService.PayOrder(orderNo, payType, time.Now())
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"code":    -1,
//			"message": "支付失败",
//			"error":   err.Error(),
//		})
//		return
//	}
//
//	// 4. 返回支付成功
//	c.JSON(http.StatusOK, gin.H{
//		"code":    0,
//		"message": "支付成功",
//	})
//}
