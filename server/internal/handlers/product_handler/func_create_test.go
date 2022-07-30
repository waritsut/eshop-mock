package product_handler

import (
	"errors"
	"eshop-mock-api/internal/models"
	mock_product_service "eshop-mock-api/internal/services/product_service/mock"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func genCreateProduct() models.Product {
	return models.Product{
		Name:              "Name Product",
		Price:             99,
		Stock:             10000,
		Catalog_Id:        1,
		Product_Status_Id: 1,
	}
}

func TestCreate(t *testing.T) {
	createProduct := genCreateProduct()
	createData := new(models.Product)
	createData.Name = createProduct.Name
	createData.Description = createProduct.Description
	createData.Price = createProduct.Price
	createData.Stock = createProduct.Stock
	createData.Catalog_Id = createProduct.Catalog_Id
	createData.Product_Status_Id = createProduct.Product_Status_Id

	testCases := []struct {
		name        string
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
		checkResponse func(t *testing.T, recoder *TestCreateContext)
	}{
		{
			name: "Created",
			bodyRequest: struct {
				name              string
				description       string
				price             int
				stock             int
				catalog_id        int
				product_status_id int
			}{
				name:              createProduct.Name,
				description:       createProduct.Description,
				price:             int(createProduct.Price),
				stock:             createProduct.Stock,
				catalog_id:        createProduct.Catalog_Id,
				product_status_id: createProduct.Product_Status_Id,
			},
			err: nil,
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Create(gomock.Eq(*createData)).
					Times(1).Return(models.Product{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *TestCreateContext) {
				require.Equal(t, http.StatusCreated, recorder.httpStatus)

			},
		},
		{
			name: "InternalError",
			bodyRequest: struct {
				name              string
				description       string
				price             int
				stock             int
				catalog_id        int
				product_status_id int
			}{
				name:              createProduct.Name,
				description:       createProduct.Description,
				price:             int(createProduct.Price),
				stock:             createProduct.Stock,
				catalog_id:        createProduct.Catalog_Id,
				product_status_id: createProduct.Product_Status_Id,
			},
			err: nil,
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Create(gomock.Eq(*createData)).
					Times(1).Return(models.Product{}, errors.New("db has problem"))
			},
			checkResponse: func(t *testing.T, recorder *TestCreateContext) {
				require.Equal(t, http.StatusInternalServerError, recorder.httpStatus)

			},
		},
		{
			name: "BadRequest",
			bodyRequest: struct {
				name              string
				description       string
				price             int
				stock             int
				catalog_id        int
				product_status_id int
			}{
				name:              "",
				description:       createProduct.Description,
				price:             0,
				stock:             createProduct.Stock,
				catalog_id:        createProduct.Catalog_Id,
				product_status_id: createProduct.Product_Status_Id,
			},
			err: errors.New("bad request"),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Create(gomock.Eq(*createData)).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *TestCreateContext) {
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
			c := &TestCreateContext{BindValue: createRequest{
				Name:              tc.bodyRequest.name,
				Description:       tc.bodyRequest.description,
				Price:             float32(tc.bodyRequest.price),
				Stock:             tc.bodyRequest.stock,
				Catalog_Id:        tc.bodyRequest.catalog_id,
				Product_Status_Id: tc.bodyRequest.product_status_id,
			}, BindError: tc.err,
			}

			productHandler.Create(c)
			tc.checkResponse(t, c)
		})
	}
}

type TestCreateContext struct {
	httpStatus int
	BindValue  createRequest
	BindError  error

	v interface{}
}

func (TestCreateContext) Method() string {
	return "POST"
}

func (TestCreateContext) Next() {
}

func (TestCreateContext) Param(v string) string {
	return ""
}

func (t TestCreateContext) Bind(v interface{}) error {
	*v.(*createRequest) = createRequest{
		Name:              t.BindValue.Name,
		Description:       t.BindValue.Description,
		Price:             t.BindValue.Price,
		Stock:             t.BindValue.Stock,
		Catalog_Id:        t.BindValue.Catalog_Id,
		Product_Status_Id: t.BindValue.Product_Status_Id,
	}
	return t.BindError
}

func (t TestCreateContext) ShouldBindQuery(v interface{}) error {
	return nil
}

func (c *TestCreateContext) JSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}

func (c *TestCreateContext) AbortWithStatusJSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}
