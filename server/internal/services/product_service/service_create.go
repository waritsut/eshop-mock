package product_service

import (
	"eshop-mock-api/internal/models"
	"time"

	"github.com/jinzhu/copier"
)

func (s *MyService) Create(createData models.Product) (model models.Product, err error) {
	copier.Copy(&model, createData)
	model.Created_At = time.Now()
	model.Updated_At = time.Now()
	if err = s.Repo.Create(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}
