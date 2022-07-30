package product_handler

import (
	"errors"
	"eshop-mock-api/internal/constants"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/response"
	"eshop-mock-api/internal/services/product_service"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
)

type DetailResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Image_Url    string `json:"imageUrl"`
	Catalog_Id   int    `json:"-"`
	Catalog_Join struct {
		Name string
	} `json:"-"`
	Catalog             string  `json:"catalog"`
	Description         string  `json:"description"`
	Price               float32 `json:"price"`
	Stock               int     `json:"stock"`
	Product_Status_Id   int     `json:"-"`
	Product_Status_Join struct {
		Name string
	} `json:"-"`
	Product_Status  string    `json:"productStatus"`
	Created_At      time.Time `json:"-"`
	Updated_At      time.Time `json:"-"`
	Created_At_Show string    `json:"createdAt"`
	Updated_At_Show string    `json:"updatedAt"`
}

func (h *Handler) Detail(c core.Context) {
	var res DetailResponse

	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		err = errors.New("bad request " + err.Error())
		response.HandleError(c, err)
		return
	}

	searchData := new(product_service.SearchDetailData)
	searchData.Id = []int{idConv}
	info, err := h.productService.Detail(*searchData,
		product_service.WithCatalog(), product_service.WithProductStatus())
	if err != nil {
		response.HandleError(c, err)
		return
	}

	copier.Copy(&res, info)
	res.prepareToDetailJson()
	response.OkWithData(c, http.StatusOK, 1, res)
}

func (d *DetailResponse) prepareToDetailJson() {
	d.Catalog = d.Catalog_Join.Name
	d.Product_Status = d.Product_Status_Join.Name
	d.Created_At_Show = d.Created_At.Format(constants.TimeLayoutDdMmYyyyHhMmSs)
	d.Updated_At_Show = d.Updated_At.Format(constants.TimeLayoutDdMmYyyyHhMmSs)
}
