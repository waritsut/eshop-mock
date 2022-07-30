package product_handler

import (
	"errors"
	mock_product_service "eshop-mock-api/internal/services/product_service/mock"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	productId := 1

	testCases := []struct {
		name          string
		productId     string
		buildStubs    func(store *mock_product_service.MockService)
		checkResponse func(t *testing.T, recoder *TestDeleteContext)
	}{
		{
			name:      "OK",
			productId: strconv.Itoa(productId),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Delete(gomock.Eq(productId)).
					Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *TestDeleteContext) {
				require.Equal(t, http.StatusOK, recorder.httpStatus)
			},
		},
		{
			name:      "InternalError",
			productId: strconv.Itoa(productId),
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Delete(gomock.Eq(productId)).
					Times(1).Return(errors.New("db has problem"))
			},
			checkResponse: func(t *testing.T, recorder *TestDeleteContext) {
				require.Equal(t, http.StatusInternalServerError, recorder.httpStatus)
			},
		},
		{
			name:      "BadRequest",
			productId: "ABC",
			buildStubs: func(store *mock_product_service.MockService) {
				store.EXPECT().Delete(gomock.Eq(productId)).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *TestDeleteContext) {
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
			c := &TestDeleteContext{paramValue: tc.productId}
			productHandler.Delete(c)
			tc.checkResponse(t, c)
		})
	}

}

type TestDeleteContext struct {
	paramValue string
	httpStatus int

	v interface{}
}

func (TestDeleteContext) Method() string {
	return "POST"
}

func (TestDeleteContext) Next() {
}

func (t TestDeleteContext) Param(v string) string {
	return t.paramValue
}

func (t TestDeleteContext) Bind(v interface{}) error {
	return nil
}

func (t TestDeleteContext) ShouldBindQuery(v interface{}) error {
	return nil
}

func (c *TestDeleteContext) JSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}

func (c *TestDeleteContext) AbortWithStatusJSON(code int, v interface{}) {
	c.httpStatus = code
	c.v = v
}
