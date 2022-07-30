package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id                int
	Name              string
	Description       string
	Image_Url         string
	Price             float32
	Stock             int
	Catalog_Id        int
	Product_Status_Id int
	Created_At        time.Time
	Updated_At        time.Time
	Deleted_At        gorm.DeletedAt
}

func (Product) TableName() string {
	return "products"
}
