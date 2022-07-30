package product_handler

import (
	"errors"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/response"
	"net/http"
)

type createRequest struct {
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description"`
	Price             float32 `json:"price" binding:"required"`
	Stock             int     `json:"stock" binding:"required"`
	Catalog_Id        int     `json:"catalog_id" binding:"required"`
	Product_Status_Id int     `json:"product_status_id" binding:"required"`
}

func (h *Handler) Create(c core.Context) {
	var req createRequest
	err := c.Bind(&req)
	if err != nil {
		err = errors.New("bad request " + err.Error())
		response.HandleError(c, err)
		return
	}

	createData := new(models.Product)
	createData.Name = req.Name
	createData.Description = req.Description
	createData.Price = req.Price
	createData.Stock = req.Stock
	createData.Catalog_Id = req.Catalog_Id
	createData.Product_Status_Id = req.Product_Status_Id

	_, err = h.productService.Create(*createData)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.OkWithDetailed(c, http.StatusCreated, 1, "success")
}
