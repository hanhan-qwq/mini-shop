package controller

import (
	"github.com/gin-gonic/gin"
	"mini_shop/model"
	"mini_shop/service"
	"mini_shop/web/request"
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

func (ctrl *ProductController) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "商品 ID 无效",
		})
		return
	}

	product, err := ctrl.ProductService.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"message": "商品不存在",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": product,
	})
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest

	// 1.绑定请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
		Status:      req.Status,
	}

	if req.CategoryID != 0 {
		product.CategoryID = &req.CategoryID
	}

	// 2.调用 service 层创建商品
	if err := ctrl.ProductService.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "创建商品失败",
			"error":   err.Error(),
		})
		return
	}

	// 3.返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建成功",
		"data":    req,
	})
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	if err := ctrl.ProductService.UpdateProduct(uint(id), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "更新失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新成功",
	})
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := ctrl.ProductService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": "删除失败", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}
func (ctrl *ProductController) AdminListProducts(c *gin.Context) {

}
