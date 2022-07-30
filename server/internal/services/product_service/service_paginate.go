package product_service

import (
	"eshop-mock-api/internal/struct_templates"
	"math"
)

type SearchPageData struct {
	Id         []int
	Catalog_Id []int

	Limit int
	Page  int
	Sort  string
}

func (s *MyService) Paginate(searchData SearchPageData, preloadOpt ...Preload) (info []Product,
	pagination struct_templates.Pagination, err error) {
	tx := s.Repo

	if len(searchData.Id) != 0 {
		tx = tx.Where("id IN (?)", searchData.Id)
	}

	if len(searchData.Catalog_Id) != 0 {
		tx = tx.Where("catalog_id IN (?)", searchData.Catalog_Id)
	}

	preload := new(PreloadModel)
	for _, f := range preloadOpt {
		f(preload)
	}

	if preload.PreloadCatalog.Flag {
		tx = tx.Preload("Catalog_Join", preload.PreloadCatalog.AddCondition...)
	}

	if preload.PreloadProductStatus.Flag {
		tx = tx.Preload("Product_Status_Join", preload.PreloadProductStatus.AddCondition...)
	}

	pagination.Limit = searchData.Limit
	pagination.Page = searchData.Page
	pagination.Sort = searchData.Sort
	tx.Model(info).Count(&pagination.TotalRows)
	pagination.TotalPages = int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
	tx.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())

	if err = tx.
		Find(&info).Error; err != nil {
		return info, pagination, err
	}
	return info, pagination, nil
}
