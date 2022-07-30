package product_handler

import (
	"errors"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/response"
	"eshop-mock-api/internal/services/product_service"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(c core.Context) {
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

	err = h.productService.Delete(info.Id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.OkWithData(c, http.StatusOK, 1, "success")
}
