package models

import (
	"database/sql"
	"time"
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
	Deleted_At        sql.NullTime
}

func (Product) TableName() string {
	return "products"
}
