package service

import (
	"mini_shop/model"
	"mini_shop/repository"
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
