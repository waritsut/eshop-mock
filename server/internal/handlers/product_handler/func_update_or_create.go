package product_handler

import (
	"errors"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/response"
	"eshop-mock-api/internal/services/product_service"
	"net/http"
	"strconv"
)

type updateOrCreateRequest struct {
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description"`
	Price             float32 `json:"price" binding:"required"`
	Stock             int     `json:"stock" binding:"required"`
	Catalog_Id        int     `json:"catalog_id" binding:"required"`
	Product_Status_Id int     `json:"product_status_id" binding:"required"`
}

func (h *Handler) UpdateOrCreate(c core.Context) {
	var req updateOrCreateRequest

	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		err = errors.New("bad request " + err.Error())
		response.HandleError(c, err)
		return
	}

	err = c.Bind(&req)
	if err != nil {
		err = errors.New("bad request " + err.Error())
		response.HandleError(c, err)
		return
	}

	whereData := new(product_service.WhereGetFirstOrCreateData)
	whereData.Id = []int{idConv}
	updateData := new(models.Product)
	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.Price = req.Price
	updateData.Stock = req.Stock
	updateData.Catalog_Id = req.Catalog_Id
	updateData.Product_Status_Id = req.Product_Status_Id
	_, err = h.productService.GetFirstOrCreateUpdate(*whereData, *updateData)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.OkWithDetailed(c, http.StatusOK, 1, "success")
}
