package response

import "eshop-mock-api/internal/core"

type Response struct {
	Id          int         `json:"id"`
	Code        string      `json:"code"`
	Description interface{} `json:"description"`
}

type ResponseData struct {
	Id   int         `json:"id"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func Message(id int, code string, description string) map[string]interface{} {
	return map[string]interface{}{
		"id":          id,
		"code":        code,
		"description": description}
}

const (
	ERROR   = "error"
	SUCCESS = "success"
)

func Result(c core.Context, status int, id int, code string, description interface{}) {
	c.JSON(status, Response{
		id,
		code,
		description,
	})
}

func ResultData(c core.Context, status int, id int, code string, data interface{}) {
	c.JSON(status, ResponseData{
		id,
		code,
		data,
	})
}

func OkWithData(c core.Context, status, id int, data interface{}) {
	ResultData(c, status, id, SUCCESS, data)
}

func OkWithDetailed(c core.Context, status, id int, description string) {
	Result(c, status, id, SUCCESS, description)
}
