package service

import (
	"errors"
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

func (s *CartService) AddToCart(userID uint, productID uint, quantity int) error {
	return s.CartDAO.AddToCart(userID, productID, quantity)
}

func (s *CartService) GetCart(userID uint) ([]CartItemResponse, float64, error) {
	items, err := s.CartDAO.GetCartItemsByUserID(userID)
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

func (s *CartService) UpdateItem(userID uint, productID uint, quantity int) error {
	cartItem, err := s.CartDAO.GetCartItem(userID, productID)
	if err != nil {
		return err
	}

	if cartItem == nil {
		return errors.New("购物车中无此商品")
	}

	cartItem.Quantity = quantity
	return s.CartDAO.UpdateCartItem(cartItem)
}

func (s *CartService) DeleteItem(userID uint, productID uint) error {
	cartItem, err := s.CartDAO.GetCartItem(userID, productID)
	if err != nil {
		return err
	}
	if cartItem == nil {
		return errors.New("购物车中无此商品")
	}

	return s.CartDAO.DeleteCartItem(cartItem)
}
