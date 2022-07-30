package product_handler

import (
	"encoding/json"
	"errors"
	"eshop-mock-api/internal/models"
	"eshop-mock-api/internal/response"
	"eshop-mock-api/internal/services/product_service"
	mock_product_service "eshop-mock-api/internal/services/product_service/mock"
	"eshop-mock-api/internal/struct_templates"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func genListProduct() []product_service.Product {
	return []product_service.Product{
		{
			Product: models.Product{
				Name:              "First Product",
				Description:       "This is good thing",
				Price:             100,
				Stock:             1000,
				Catalog_Id:        1,
				Product_Status_Id: 1,
				Created_At:        time.Now(),
				Updated_At:        time.Now(),
				Deleted_At:        gorm.DeletedAt{},
			},
			Catalog_Join:        models.Catalog{Id: 1, Name: "Furniture"},
			Product_Status_Join: models.ProductStatus{Id: 1, Name: "On Shelf"},
		},
		{
			Product: models.Product{
				Name:              "Second Product",
				Description:       "This is bad thing",
				Price:             10,
				Stock:             100,
				Catalog_Id:        1,
				Product_Status_Id: 1,
				Created_At:        time.Now(),
				Updated_At:        time.Now(),
				Deleted_At:        gorm.DeletedAt{},
			},
			Catalog_Join:        models.Catalog{Id: 1, Name: "Furniture"},
			Product_Status_Join: models.ProductStatus{Id: 1, Name: "On Shelf"},
		},
	}
}

func genPagination() struct_templates.Pagination {
	return struct_templates.Pagination{
		Limit:      2,
		Page:       1,
		Sort:       "created_at DESC",
		TotalRows:  10,
		TotalPages: 5,
		PrevPage:   1,
		NextPage:   1,
	}
}

func TestPaginate(t *testing.T) {
	productList := genListProduct()
	pagination := genPagination()
	CatalogId := []int{1}
	searchPage := new(product_service.SearchPageData)
	searchPage.Catalog_Id = CatalogId
	searchPage.Limit = pagination.Limit
	searchPage.Page = pagination.Page
	searchPage.Sort = pagination.Sort

	testCases := []struct {
		name          string
		catalogId     []int
		limit         int
		page          int
		sort          string
		err           error
		buildStubs    func(store *mock_product_service.MockService)
		checkResponse func(t *testing.T, recoder *TestPagContext)
	}{
		{
			name:      "OK",
			catalogId: CatalogId,
			limit:     pagination.Limit,
			page:      pagination.Page,
			sort:      searchPage.Sort,
			err:       nil,
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Paginate(gomock.Eq(*searchPage), gomock.Any(), gomock.Any()).
					Times(1).Return(productList, pagination, nil)
			},
			checkResponse: func(t *testing.T, recorder *TestPagContext) {
				require.Equal(t, http.StatusOK, recorder.httpStatus)
				requireBodyMatchProductPaginate(t, recorder, productList, pagination)
			},
		},
		{
			name:      "InternalError",
			catalogId: CatalogId,
			limit:     pagination.Limit,
			page:      pagination.Page,
			sort:      searchPage.Sort,
			err:       nil,
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Paginate(gomock.Eq(*searchPage), gomock.Any(), gomock.Any()).
					Times(1).Return(productList, pagination, errors.New("db has problem"))
			},
			checkResponse: func(t *testing.T, recorder *TestPagContext) {
				var res response.ResponseData
				bodyBytes, _ := json.Marshal(recorder.v)
				json.Unmarshal(bodyBytes, &res)
				require.Equal(t, http.StatusInternalServerError, recorder.httpStatus)
			},
		},
		{
			name:      "BadRequest",
			catalogId: []int{},
			limit:     0,
			page:      0,
			sort:      searchPage.Sort,
			err:       errors.New("bad request"),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Paginate(gomock.Eq(*searchPage), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *TestPagContext) {
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
			c := &TestPagContext{
				ShouldBindQueryValue: ListRequest{
					Catalog_Id: tc.catalogId,
					Limit:      tc.limit,
					Page:       tc.page,
					Sort:       tc.sort,
				},
				ShouldBindQueryError: tc.err,
			}

			productHandler.Paginate(c)
			tc.checkResponse(t, c)

		})
	}
}

func requireBodyMatchProductPaginate(t *testing.T, body *TestPagContext,
	product []product_service.Product, pagination struct_templates.Pagination) {
	var gotRes response.ResponseData
	bodyBytes, err := json.Marshal(body.v)
	require.NoError(t, err)
	json.Unmarshal(bodyBytes, &gotRes)
	gotJson := gotRes.Data.(map[string]interface{})

	expectRes := ListResponse{}
	copier.Copy(&expectRes.Rows, product)
	copier.Copy(&expectRes, pagination)
	expectRes.prepareToPageJson()
	MarshalOut, _ := json.MarshalIndent(expectRes, "", "  ")
	expectJson := make(map[string]interface{})
	json.Unmarshal(MarshalOut, &expectJson)

	require.Equal(t, true, reflect.DeepEqual(gotJson, expectJson))
}

type TestPagContext struct {
	httpStatus           int
	ShouldBindQueryValue ListRequest
	ShouldBindQueryError error

	v interface{}
}

func (TestPagContext) Method() string {
	return "GET"
}

func (TestPagContext) Next() {
}

func (t TestPagContext) Param(v string) string {
	return ""
}

func (TestPagContext) Bind(v interface{}) error {
	return nil
}

func (t TestPagContext) ShouldBindQuery(v interface{}) error {
	*v.(*ListRequest) = ListRequest{
		Catalog_Id: t.ShouldBindQueryValue.Catalog_Id,
		Limit:      t.ShouldBindQueryValue.Limit,
		Page:       t.ShouldBindQueryValue.Page,
		Sort:       t.ShouldBindQueryValue.Sort,
	}
	return t.ShouldBindQueryError
}

func (c *TestPagContext) JSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}

func (c *TestPagContext) AbortWithStatusJSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}
