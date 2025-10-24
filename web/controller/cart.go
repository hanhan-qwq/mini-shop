package controller

import (
	"mini_shop/service"
)

type CartController struct {
	CartService *service.CartService
}

func NewCartController() *CartController {
	return &CartController{
		CartService: service.NewCartService(),
	}
}
