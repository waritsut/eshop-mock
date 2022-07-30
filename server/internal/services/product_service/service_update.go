package product_service

import (
	"eshop-mock-api/internal/models"
	"time"
)

type WhereUpdateData struct {
	Id []int
}

func (s *MyService) Update(selectField []string, whereData WhereUpdateData, updateData models.Product) (err error) {
	tx := s.Repo

	if len(whereData.Id) != 0 {
		tx = tx.Where("id IN (?)", whereData.Id)
	}

	if err = tx.Table("products").
		Select(selectField, "updated_at").
		Updates(map[string]interface{}{
			"name":              updateData.Name,
			"description":       updateData.Description,
			"price":             updateData.Price,
			"stock":             updateData.Stock,
			"catalog_id":        updateData.Catalog_Id,
			"product_status_id": updateData.Product_Status_Id,
			"updated_at":        time.Now(),
		}).Error; err != nil {
		return err
	}
	return

}
