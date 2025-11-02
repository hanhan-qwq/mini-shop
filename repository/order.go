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

// FindUserOrders 查询用户订单（分页 + 状态过滤）
func (d *OrderDAO) FindUserOrders(userID uint, page, pageSize, status int) ([]model.Order, int64, error) {
	var (
		orders []model.Order
		count  int64
	)

	query := d.db.Model(&model.Order{}).Where("user_id = ?", userID)

	// 状态过滤（status=0 表示全部）
	if status != 0 {
		query = query.Where("status = ?", status)
	}

	// 统计总数
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 查询分页数据（按创建时间倒序）
	offset := (page - 1) * pageSize
	err := query.
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&orders).Error

	if err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}
