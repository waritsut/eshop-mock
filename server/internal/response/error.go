package response

import (
	"eshop-mock-api/internal/core"
	"net/http"
	"strings"
)

func HandleError(c core.Context, err error) {
	if strings.Contains(err.Error(), "bad request") {
		Result(c, http.StatusBadRequest, -1, ERROR, err.Error())
	} else if strings.Contains(err.Error(), "record not found") {
		Result(c, http.StatusNotFound, -1, ERROR, err.Error())
	} else {
		Result(c, http.StatusInternalServerError, -1, ERROR, err.Error())
	}
}

func AbortWithDetail(c core.Context, status, id int, description string) {
	c.AbortWithStatusJSON(status, Response{
		id,
		ERROR,
		description,
	})
}
