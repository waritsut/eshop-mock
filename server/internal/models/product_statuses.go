package models

type ProductStatus struct {
	Id   int
	Name string
}

func (ProductStatus) TableName() string {
	return "product_statuses"
}
