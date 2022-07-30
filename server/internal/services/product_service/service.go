package product_service

import (
	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/store"
	"eshop-mock-api/internal/struct_templates"
)

type Service interface {
	Create(createData models.Product) (model models.Product, err error)
	Delete(id int) error
	Detail(searchData SearchDetailData, preloadOpt ...Preload) (info Product, err error)
	GetFirstOrCreateUpdate(whereData WhereGetFirstOrCreateData, attrsData models.Product) (info models.Product, err error)
	Paginate(searchData SearchPageData, preloadOpt ...Preload) (info []Product,
		pagination struct_templates.Pagination, err error)
	Update(selectField []string, whereData WhereUpdateData, updateData models.Product) (err error)
}

type MyService struct {
	store.Repo
}

func New(db store.Repo) Service {
	return &MyService{
		Repo: db,
	}
}
