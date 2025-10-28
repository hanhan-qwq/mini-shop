package repository

import (
	"gorm.io/gorm"
	"mini_shop/global"
	"mini_shop/model"
)

type CartDAO struct {
	db *gorm.DB
}

func NewCartDAO() *CartDAO {
	return &CartDAO{
		db: global.GetDB(),
	}
}

func (d *CartDAO) CreateItem(item *model.CartItem) error {
	return d.db.Create(item).Error
}

func (d *CartDAO) GetCartItemsByUserID(userID uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	err := d.db.Where("user_id = ?", userID).Find(&cartItems).Error
	return cartItems, err
}

func (d *CartDAO) GetCartItem(userID uint, productID uint) (*model.CartItem, error) {
	var cartItem model.CartItem
	err := d.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func (d *CartDAO) UpdateCartItem(item *model.CartItem) error {
	return d.db.Save(item).Error
}

func (d *CartDAO) DeleteCartItem(item *model.CartItem) error {
	return d.db.Delete(item).Error
}
