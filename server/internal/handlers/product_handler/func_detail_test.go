package product_handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/response"
	"eshop-mock-api/internal/services/product_service"
	mock_product_service "eshop-mock-api/internal/services/product_service/mock"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

func genDetailProduct() product_service.Product {
	return product_service.Product{
		Product: models.Product{
			Id:                1,
			Name:              "Phone",
			Description:       "Good Phone",
			Price:             1000,
			Stock:             100,
			Catalog_Id:        1,
			Product_Status_Id: 1,
			Created_At:        time.Now(),
			Updated_At:        time.Now(),
		},
		Catalog_Join: models.Catalog{
			Id:   2,
			Name: "Electric Appliance",
		},
		Product_Status_Join: models.ProductStatus{
			Id:   1,
			Name: "On Shelf",
		},
	}
}

func TestDetail(t *testing.T) {
	product := genDetailProduct()
	searchData := new(product_service.SearchDetailData)
	searchData.Id = []int{product.Id}

	testCases := []struct {
		name          string
		productId     string
		buildStubs    func(store *mock_product_service.MockService)
		checkResponse func(t *testing.T, recoder *TestDetailContext)
	}{
		{
			name:      "OK",
			productId: strconv.Itoa(product.Id),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Detail(gomock.Eq(*searchData), gomock.Any(), gomock.Any()).
					Times(1).Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *TestDetailContext) {
				require.Equal(t, http.StatusOK, recorder.httpStatus)
				requireBodyMatchProduct(t, recorder, product)
			},
		},
		{
			name:      " NotFound",
			productId: strconv.Itoa(product.Id),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Detail(gomock.Eq(*searchData), gomock.Any(), gomock.Any()).
					Times(1).Return(product_service.Product{}, errors.New("record not found"))
			},
			checkResponse: func(t *testing.T, recorder *TestDetailContext) {
				var res response.ResponseData
				bodyBytes, _ := json.Marshal(recorder.v)
				json.Unmarshal(bodyBytes, &res)
				require.Equal(t, http.StatusNotFound, recorder.httpStatus)
			},
		},
		{
			name:      "InternalError",
			productId: strconv.Itoa(product.Id),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Detail(gomock.Eq(*searchData), gomock.Any(), gomock.Any()).
					Times(1).Return(product_service.Product{}, errors.New("db has problem"))
			},
			checkResponse: func(t *testing.T, recorder *TestDetailContext) {
				var res response.ResponseData
				bodyBytes, _ := json.Marshal(recorder.v)
				json.Unmarshal(bodyBytes, &res)
				require.Equal(t, http.StatusInternalServerError, recorder.httpStatus)
			},
		},
		{
			name:      "BadRequest",
			productId: "ABC",
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Detail(gomock.Eq(*searchData), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *TestDetailContext) {
				var res response.ResponseData
				bodyBytes, _ := json.Marshal(recorder.v)
				json.Unmarshal(bodyBytes, &res)
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
			c := &TestDetailContext{paramValue: tc.productId}

			productHandler.Detail(c)
			tc.checkResponse(t, c)
		})
	}
}

func requireBodyMatchProduct(t *testing.T, body *TestDetailContext, product product_service.Product) {
	var gotRes response.ResponseData
	bodyBytes, err := json.Marshal(body.v)
	require.NoError(t, err)
	json.Unmarshal(bodyBytes, &gotRes)
	gotJson := gotRes.Data.(map[string]interface{})

	expectRes := DetailResponse{}
	copier.Copy(&expectRes, product)
	expectRes.prepareToDetailJson()
	MarshalOut, _ := json.MarshalIndent(expectRes, "", "  ")
	expectJson := make(map[string]interface{})
	json.Unmarshal(MarshalOut, &expectJson)

	require.Equal(t, true, reflect.DeepEqual(gotJson, expectJson))
}

type TestDetailContext struct {
	paramValue string
	httpStatus int

	v interface{}
}

func (TestDetailContext) Method() string {
	return "GET"
}

func (TestDetailContext) Next() {
}

func (t TestDetailContext) Param(v string) string {
	return t.paramValue
}

func (TestDetailContext) Bind(v interface{}) error {
	return nil
}

func (TestDetailContext) ShouldBindQuery(v interface{}) error {
	return nil
}

func (c *TestDetailContext) JSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}

func (c *TestDetailContext) AbortWithStatusJSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}
