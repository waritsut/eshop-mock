package product_handler

import (
	"eshop-mock-api/internal/services/product_service"
)

type Handler struct {
	productService product_service.Service
}

func New(productService product_service.Service) *Handler {
	return &Handler{
		productService: productService,
	}
}
