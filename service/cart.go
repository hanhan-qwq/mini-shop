package service

import (
	"mini_shop/repository"
)

type CartService struct {
	CartDAO    *repository.CartDAO
	ProductDAO *repository.ProductDAO
}

func NewCartService() *CartService {
	return &CartService{
		CartDAO:    repository.NewCartDAO(),
		ProductDAO: repository.NewProductDAO(),
	}
}

type CartItemResponse struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	ImageURL  string  `json:"image_url"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
	Checked   bool    `json:"checked"`
}

func (s *CartService) AddToCart(userId uint, productId uint, quantity int) error {
	return s.CartDAO.AddToCart(userId, productId, quantity)
}

func (s *CartService) GetCart(userId uint) ([]CartItemResponse, float64, error) {
	items, err := s.CartDAO.GetCartItemsByUserID(userId)
	if err != nil {
		return nil, 0, err
	}

	var cartList []CartItemResponse
	var total float64

	for _, item := range items {
		product, err := s.ProductDAO.GetProductByID(item.ProductID)
		if err != nil {
			return nil, 0, err
		}

		subtotal := product.Price * float64(item.Quantity)

		if item.Checked {
			total += subtotal
		}

		cartList = append(cartList, CartItemResponse{
			ProductID: product.ID,
			Name:      product.Name,
			ImageURL:  product.ImageURL,
			Price:     subtotal,
			Quantity:  item.Quantity,
			Subtotal:  subtotal,
			Checked:   item.Checked,
		})
	}

	return cartList, total, nil
}
