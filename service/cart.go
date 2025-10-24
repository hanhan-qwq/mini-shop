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
