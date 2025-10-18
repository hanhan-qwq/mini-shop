package repository

import (
	"gorm.io/gorm"
	"mini_shop/global"
	"mini_shop/model"
)

type ProductDAO struct {
	db *gorm.DB
}

func NewProductDAO() *ProductDAO {
	return &ProductDAO{
		db: global.GetDB(),
	}
}

func (d *ProductDAO) ListProducts(offset, pageSize, categoryID int, keyword, sort string) ([]model.Product, int64, error) {
	var products []model.Product
	var count int64

	currentDB := d.db.Model(&model.Product{}).Where("status = ?", "on_sale")

	if keyword != "" {
		currentDB = currentDB.Where("`name` LIKE ?", "%"+keyword+"%")
	}
	if categoryID != 0 {
		currentDB = currentDB.Where("`category_id` = ?", categoryID)
	}

	switch sort {
	case "price_desc":
		currentDB = currentDB.Order("price DESC")
	case "price_asc":
		currentDB = currentDB.Order("price ASC")
	case "sold_desc":
		currentDB = currentDB.Order("sold DESC")
	default:
		currentDB = currentDB.Order("created_at DESC")
	}

	currentDB.Count(&count)
	err := currentDB.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}
