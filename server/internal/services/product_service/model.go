package product_service

import (
	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/struct_templates"
)

type Product struct {
	models.Product

	Catalog_Join        models.Catalog       `gorm:"foreignKey:id;references:catalog_id"`
	Product_Status_Join models.ProductStatus `gorm:"foreignKey:id;references:product_status_id"`
}

type Preload func(*PreloadModel)

type PreloadModel struct {
	PreloadCatalog       struct_templates.PreloadOption
	PreloadProductStatus struct_templates.PreloadOption
}

func WithCatalog(con ...interface{}) Preload {
	return func(pre *PreloadModel) {
		pre.PreloadCatalog.Flag = true
		pre.PreloadCatalog.AddCondition = con
	}
}

func WithProductStatus(con ...interface{}) Preload {
	return func(pre *PreloadModel) {
		pre.PreloadProductStatus.Flag = true
		pre.PreloadProductStatus.AddCondition = con
	}
}
