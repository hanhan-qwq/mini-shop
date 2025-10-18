package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`           // 商品名称
	Description string  `gorm:"type:text" json:"description"`                     // 商品描述
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`         // 单价
	Stock       int     `gorm:"not null;default:0" json:"stock"`                  // 库存数量
	ImageURL    string  `gorm:"type:varchar(255)" json:"image_url"`               // 商品主图
	Status      string  `gorm:"type:varchar(20);default:'on_sale'" json:"status"` // 商品状态: on_sale=上架, =下架
	CategoryID  *uint   `json:"category_id"`                                      // 分类ID（可以为空）
}
