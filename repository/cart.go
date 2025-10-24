package repository

import (
	"gorm.io/gorm"
	"mini_shop/global"
)

type CartDAO struct {
	db *gorm.DB
}

func NewCartDAO() *CartDAO {
	return &CartDAO{
		db: global.GetDB(),
	}
}
