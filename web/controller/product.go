package controller

import (
	"github.com/gin-gonic/gin"
	"mini_shop/service"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		ProductService: service.NewProductService(),
	}
}

func (ctrl *ProductController) ListProducts(c *gin.Context) {
	//1.获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	sort := c.Query("sort")
	categoryID, _ := strconv.Atoi(c.DefaultQuery("category_id", "0"))
	//2.按条件查询
	products, count, err := ctrl.ProductService.ListProducts(page, pageSize, categoryID, keyword, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "查询产品数据失败",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"product":   products,
		"count":     count,
		"page":      page,
		"page_size": pageSize,
	})
}
