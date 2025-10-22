package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID    uint `gorm:"not null;index" json:"user_id"`
	ProductID uint `gorm:"not null;index" json:"product_id"`
	Quantity  int  `gorm:"not null;default:1" json:"quantity"`
	Checked   bool `gorm:"default:true" json:"checked"`
}
