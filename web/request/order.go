package request

type CreateOrderRequest struct {
	Remark string                   `json:"remark"` // 可选
	Items  []CreateOrderItemRequest `json:"items" binding:"required"`
}

type CreateOrderItemRequest struct {
	ProductID   uint    `json:"product_id" binding:"required"`
	ProductName string  `json:"product_name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
}
