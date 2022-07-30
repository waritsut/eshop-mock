package product_service

import "eshop-mock-api/internal/models"

func (s *MyService) Delete(id int) error {
	model := models.Product{Id: id}
	if err := s.Repo.Delete(&model).Error; err != nil {
		return err
	}
	return nil
}
