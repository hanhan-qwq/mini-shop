package repository

import (
	"gorm.io/gorm"
	"mini_shop/global"
	"mini_shop/model"
)

type OrderDAO struct {
	db *gorm.DB
}

func NewOrderDAO() *OrderDAO {
	return &OrderDAO{
		db: global.GetDB(),
	}
}

func (d *OrderDAO) CreateOrderInTx(tx *gorm.DB, order *model.Order, items []model.OrderItem) error {
	if err := tx.Create(order).Error; err != nil {
		return err
	}
	for i := range items {
		items[i].OrderID = order.ID
	}
	if err := tx.Create(&items).Error; err != nil {
		return err
	}
	return nil
}
