package product_handler

import (
	"errors"
	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/services/product_service"
	mock_product_service "eshop-mock-api/internal/services/product_service/mock"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func genUpdateOrCreateProduct() models.Product {
	return models.Product{
		Name:              "Name Product",
		Price:             99,
		Stock:             10000,
		Catalog_Id:        1,
		Product_Status_Id: 1,
	}
}

func genWhereUpdateOrCreateProduct() product_service.WhereGetFirstOrCreateData {
	return product_service.WhereGetFirstOrCreateData{
		Id: []int{1},
	}
}

func TestUpdateOrCreate(t *testing.T) {
	updateProduct := genUpdateOrCreateProduct()
	updateData := new(models.Product)
	updateData.Name = updateProduct.Name
	updateData.Description = updateProduct.Description
	updateData.Price = updateProduct.Price
	updateData.Stock = updateProduct.Stock
	updateData.Catalog_Id = updateProduct.Catalog_Id
	updateData.Product_Status_Id = updateProduct.Product_Status_Id

	whereProduct := genWhereUpdateOrCreateProduct()
	whereData := new(product_service.WhereGetFirstOrCreateData)
	whereData.Id = whereProduct.Id

	testCases := []struct {
		name        string
		productId   string
		bodyRequest struct {
			name              string
			description       string
			price             int
			stock             int
			catalog_id        int
			product_status_id int
		}
		err           error
		buildStubs    func(store *mock_product_service.MockService)
		checkResponse func(t *testing.T, recoder *TestUpdateOrCreateContext)
	}{
		{
			name:      "OK",
			productId: strconv.Itoa(whereData.Id[0]),
			bodyRequest: struct {
				name              string
				description       string
				price             int
				stock             int
				catalog_id        int
				product_status_id int
			}{
				name:              updateData.Name,
				description:       updateData.Description,
				price:             int(updateData.Price),
				stock:             updateData.Stock,
				catalog_id:        updateData.Catalog_Id,
				product_status_id: updateData.Product_Status_Id,
			},
			err: nil,
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().GetFirstOrCreateUpdate(gomock.Eq(*whereData), gomock.Eq(*updateData)).
					Times(1).Return(models.Product{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *TestUpdateOrCreateContext) {
				require.Equal(t, http.StatusOK, recorder.httpStatus)

			},
		},
		{
			name:      "InternalError",
			productId: strconv.Itoa(whereData.Id[0]),
			bodyRequest: struct {
				name              string
				description       string
				price             int
				stock             int
				catalog_id        int
				product_status_id int
			}{
				name:              updateData.Name,
				description:       updateData.Description,
				price:             int(updateData.Price),
				stock:             updateData.Stock,
				catalog_id:        updateData.Catalog_Id,
				product_status_id: updateData.Product_Status_Id,
			},
			err: nil,
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().GetFirstOrCreateUpdate(gomock.Eq(*whereData), gomock.Eq(*updateData)).
					Times(1).Return(models.Product{}, errors.New("db has problem"))
			},
			checkResponse: func(t *testing.T, recorder *TestUpdateOrCreateContext) {
				require.Equal(t, http.StatusInternalServerError, recorder.httpStatus)
			},
		},
		{
			name:      "BadRequest",
			productId: "",
			bodyRequest: struct {
				name              string
				description       string
				price             int
				stock             int
				catalog_id        int
				product_status_id int
			}{
				name:              updateData.Name,
				description:       updateData.Description,
				price:             int(updateData.Price),
				stock:             updateData.Stock,
				catalog_id:        updateData.Catalog_Id,
				product_status_id: updateData.Product_Status_Id,
			},
			err: errors.New("bad request"),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().GetFirstOrCreateUpdate(gomock.Eq(*whereData), gomock.Eq(*updateData)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *TestUpdateOrCreateContext) {
				require.Equal(t, http.StatusBadRequest, recorder.httpStatus)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_product_service.NewMockService(ctrl)
			tc.buildStubs(store)

			productHandler := New(store)
			c := &TestUpdateOrCreateContext{
				paramValue: tc.productId,
				BindValue: updateOrCreateRequest{
					Name:              tc.bodyRequest.name,
					Description:       tc.bodyRequest.description,
					Price:             float32(tc.bodyRequest.price),
					Stock:             tc.bodyRequest.stock,
					Catalog_Id:        tc.bodyRequest.catalog_id,
					Product_Status_Id: tc.bodyRequest.product_status_id,
				}, BindError: tc.err,
			}

			productHandler.UpdateOrCreate(c)
			tc.checkResponse(t, c)
		})
	}

}

type TestUpdateOrCreateContext struct {
	httpStatus int
	paramValue string
	BindValue  updateOrCreateRequest
	BindError  error

	v interface{}
}

func (TestUpdateOrCreateContext) Method() string {
	return "PUT"
}

func (TestUpdateOrCreateContext) Next() {
}

func (t TestUpdateOrCreateContext) Param(v string) string {
	return t.paramValue
}

func (t TestUpdateOrCreateContext) Bind(v interface{}) error {
	*v.(*updateOrCreateRequest) = updateOrCreateRequest{
		Name:              t.BindValue.Name,
		Description:       t.BindValue.Description,
		Price:             t.BindValue.Price,
		Stock:             t.BindValue.Stock,
		Catalog_Id:        t.BindValue.Catalog_Id,
		Product_Status_Id: t.BindValue.Product_Status_Id,
	}
	return t.BindError
}

func (t TestUpdateOrCreateContext) ShouldBindQuery(v interface{}) error {
	return nil
}

func (c *TestUpdateOrCreateContext) JSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}

func (c *TestUpdateOrCreateContext) AbortWithStatusJSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}
