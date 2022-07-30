package product_service

type SearchDetailData struct {
	Id []int
}

func (s *MyService) Detail(searchData SearchDetailData, preloadOpt ...Preload) (info Product, err error) {
	tx := s.Repo

	if len(searchData.Id) != 0 {
		tx = tx.Where("id IN (?)", searchData.Id)
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

	if err = tx.
		First(&info).Error; err != nil {
		return info, err
	}

	return info, nil
}
