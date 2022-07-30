package product_handler

import (
	"errors"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/response"
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

	err = h.productService.Delete(idConv)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.OkWithData(c, http.StatusOK, 1, "success")
}
