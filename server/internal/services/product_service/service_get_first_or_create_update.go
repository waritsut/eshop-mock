package product_service

import (
	"eshop-mock-api/internal/models"
	"time"
)

type WhereGetFirstOrCreateData struct {
	Id []int
}

func (s *MyService) GetFirstOrCreateUpdate(whereData WhereGetFirstOrCreateData, attrsData models.Product) (info models.Product, err error) {
	tx := s.Repo

	if len(whereData.Id) != 0 {
		tx = tx.Where("id IN (?)", whereData.Id)
	}

	attrsData.Created_At = time.Now()
	attrsData.Updated_At = time.Now()
	tx.
		Unscoped().
		Attrs(attrsData).
		FirstOrCreate(&info)

	if err := s.Repo.Model(&info).
		Unscoped().
		Updates(map[string]interface{}{
			"name":              attrsData.Name,
			"description":       attrsData.Description,
			"price":             attrsData.Price,
			"stock":             attrsData.Stock,
			"catalog_id":        attrsData.Catalog_Id,
			"product_status_id": attrsData.Product_Status_Id,
			"updated_at":        time.Now(),
			"deleted_at":        nil,
		}).Error; err != nil {
		return info, err
	}

	return
}
