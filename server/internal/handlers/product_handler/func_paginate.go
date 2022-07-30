package product_handler

import (
	"errors"
	"eshop-mock-api/internal/constants"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/response"
	"eshop-mock-api/internal/services/product_service"
	"eshop-mock-api/internal/struct_templates"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
)

type ListRequest struct {
	Catalog_Id []int  `form:"catalog_id[]" binding:"required"`
	Limit      int    `form:"limit" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	Sort       string `form:"sort,default=created_at DESC"`
}

type ListResponse struct {
	Rows []struct {
		Id         int         `json:"id"`
		Name       string      `json:"name"`
		Image_Url  string      `json:"imageUrl"`
		Price      float32     `json:"price"`
		Stock      int         `json:"stock"`
		Created_At interface{} `json:"createdAt"`
		Updated_At interface{} `json:"updatedAt"`
	} `json:"rows"`

	struct_templates.Pagination
}

func (h *Handler) Paginate(c core.Context) {
	var req ListRequest
	var res ListResponse

	if err := c.ShouldBindQuery(&req); err != nil {
		err = errors.New("bad request " + err.Error())
		response.HandleError(c, err)
		return
	}

	searchPage := new(product_service.SearchPageData)
	searchPage.Catalog_Id = req.Catalog_Id
	searchPage.Limit = req.Limit
	searchPage.Page = req.Page
	searchPage.Sort = req.Sort
	list, pagination, err := h.productService.Paginate(*searchPage,
		product_service.WithCatalog(),
		product_service.WithProductStatus())
	if err != nil {
		response.HandleError(c, err)
		return
	}

	copier.Copy(&res, pagination)
	copier.Copy(&res.Rows, list)
	res.prepareToPageJson()
	response.OkWithData(c, http.StatusOK, 1, res)
}

func (l *ListResponse) prepareToPageJson() {
	for index := range l.Rows {
		l.Rows[index].Created_At = l.Rows[index].Created_At.(time.Time).Format(constants.TimeLayoutDdMmYyyyHhMmSs)
		l.Rows[index].Updated_At = l.Rows[index].Updated_At.(time.Time).Format(constants.TimeLayoutDdMmYyyyHhMmSs)

	}
}
