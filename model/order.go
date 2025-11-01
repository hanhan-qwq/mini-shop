package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderNo    string     `gorm:"unique;not null;size:64" json:"order_no"` // 订单编号（业务唯一，如时间戳+随机数）
	UserID     uint       `gorm:"not null;index" json:"user_id"`           // 关联用户ID
	TotalPrice float64    `gorm:"not null" json:"total_price"`             // 订单总金额（元）
	Status     int        `gorm:"not null" json:"status"`                  // 订单状态（核心流转字段）
	PayTime    *time.Time `json:"pay_time"`                                // 支付时间（可为空，未支付时为nil）
	//AddressID  uint64         `gorm:"not null" json:"address_id"`           // 收货地址ID（关联地址表）
	Remark string `gorm:"size:500" json:"remark"` // 订单备注（用户填写）
}

// OrderItem 订单项
type OrderItem struct {
	gorm.Model
	OrderID     uint    `gorm:"not null;index" json:"order_id"` // 关联订单ID
	ProductID   uint    `gorm:"not null" json:"product_id"`     // 商品ID
	ProductName string  `gorm:"not null" json:"product_name"`   // 商品名称（快照，下单时的名称）
	Price       float64 `gorm:"not null" json:"price"`          // 下单时的单价（快照）
	Quantity    int     `gorm:"not null" json:"quantity"`       // 购买数量
	TotalPrice  float64 `gorm:"not null" json:"total_price"`    // 小计金额（单价×数量）
}

const (
	OrderStatusPendingPay = 10 // 待支付（刚创建的订单）
	OrderStatusPaid       = 20 // 已支付（用户付款后）
	OrderStatusCancelled  = 30 // 已取消（用户未支付时取消）
)
