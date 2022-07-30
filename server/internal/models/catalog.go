package models

type Catalog struct {
	Id   int
	Name string
}

func (Catalog) TableName() string {
	return "catalogs"
}
