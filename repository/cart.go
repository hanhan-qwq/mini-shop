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

func (d *CartDAO) AddToCart(userId uint, productId uint, quantity int) error {
	var cartItem model.CartItem

	err := d.db.Where("user_id = ? AND product_id = ?", userId, productId).Find(&cartItem).Error
	if err == nil {
		cartItem.Quantity += quantity
		return d.db.Save(&cartItem).Error
	}

	cartItem = model.CartItem{
		UserID:    userId,
		ProductID: productId,
		Quantity:  quantity,
	}

	return d.db.Create(&cartItem).Error
}

func (d *CartDAO) GetCartItemsByUserID(userId uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	err := d.db.Where("user_id = ?", userId).Find(&cartItems).Error
	return cartItems, err
}

func (d *CartDAO) GetCartItem(userID uint, productID uint) (*model.CartItem, error) {
	var cartItem model.CartItem
	err := d.db.Where("user_id = ? AND product_id = ?", userID, productID).Find(&cartItem).Error
	return &cartItem, err
}

func (d *CartDAO) UpdateCartItem(item *model.CartItem) error {
	return d.db.Save(item).Error
}
