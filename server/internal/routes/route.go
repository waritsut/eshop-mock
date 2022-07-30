package routes

import (
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/handlers/product_handler"
	"eshop-mock-api/internal/services"
)

type Handlers struct {
	productCtrl *product_handler.Handler
}

func New(ServiceCollection services.ServiceCollection) *Handlers {
	return &Handlers{
		productCtrl: product_handler.New(ServiceCollection.ProductService),
	}
}

func (h *Handlers) RegisterRoutes(r core.Mux) *Handlers {
	apiGroup := r.Group("/api")
	productGroup := apiGroup.Group("products")
	{
		productGroup.GET("", h.productCtrl.Paginate)
		productGroup.GET("/:id", h.productCtrl.Detail)
		productGroup.POST("", h.productCtrl.Create)
		productGroup.PUT("/:id", h.productCtrl.UpdateOrCreate)
		productGroup.DELETE("/:id")
	}
	return h
}
