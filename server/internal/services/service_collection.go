package services

import (
	"eshop-mock-api/internal/services/product_service"
	"eshop-mock-api/internal/store"
)

type ServiceCollection struct {
	ProductService product_service.Service
}

func New(db store.Repo) ServiceCollection {
	return ServiceCollection{
		ProductService: product_service.New(db),
	}
}
