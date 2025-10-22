package service

import (
	"errors"
	"fmt"
	"mini_shop/model"
	"mini_shop/repository"
	"mini_shop/web/request"
)

type ProductService struct {
	ProductDB *repository.ProductDAO
}

func NewProductService() *ProductService {
	return &ProductService{
		ProductDB: repository.NewProductDAO(),
	}
}

func (s *ProductService) ListProducts(page, pageSize, categoryID int, keyword, sort string) ([]model.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return s.ProductDB.ListProducts(offset, pageSize, categoryID, keyword, sort)
}

func (s *ProductService) GetProduct(id uint) (*model.Product, error) {
	product, err := s.ProductDB.GetProductByID(id)
	if err != nil {
		return nil, errors.New("商品不存在或已下架")
	}
	return product, nil
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.ProductDB.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(id uint, req *request.UpdateProductRequest) error {
	// 1.查询商品是否存在
	product, err := s.ProductDB.GetProductByID(id)
	if err != nil {
		return fmt.Errorf("商品不存在: %w", err)
	}

	// 2.更新字段
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 { // 可以为0
		product.Stock = req.Stock
	}
	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}
	if req.Status != "" {
		product.Status = req.Status
	}
	if req.CategoryID != 0 {
		product.CategoryID = &req.CategoryID
	}

	return s.ProductDB.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.ProductDB.DeleteProduct(id)
}
