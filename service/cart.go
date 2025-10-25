package service

import "mini_shop/repository"

type CartService struct {
	CartDAO *repository.CartDAO
}

func NewCartService() *CartService {
	return &CartService{
		CartDAO: repository.NewCartDAO(),
	}
}

func (service *CartService) AddToCart(userId uint, productId uint, quantity int) error {
	return service.CartDAO.AddToCart(userId, productId, quantity)
}
