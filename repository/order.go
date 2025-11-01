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

func (d *OrderDAO) GetOrderByID(userID, orderID uint) (*model.Order, error) {
	var order model.Order
	if err := d.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (d *OrderDAO) GetOrderItems(orderID uint) ([]model.OrderItem, error) {
	var items []model.OrderItem
	if err := d.db.Where("order_id = ?", orderID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
