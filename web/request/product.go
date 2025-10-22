package request

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,max=100"`
	Description string  `json:"description" binding:"omitempty,max=500"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	ImageURL    string  `json:"image_url" binding:"omitempty"`
	CategoryID  uint    `json:"category_id" binding:"omitempty"`
	Status      string  `json:"status" binding:"omitempty,oneof=on_sale off_sale"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required,max=100"`
	Description string  `json:"description" binding:"omitempty,max=500"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	ImageURL    string  `json:"image_url" binding:"omitempty"`
	CategoryID  uint    `json:"category_id" binding:"omitempty"`
	Status      string  `json:"status" binding:"omitempty,oneof=on_sale off_sale"`
}
