package handlers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"quotetion_book/internal/delivery/http/handlers"
	"quotetion_book/internal/models"
	"quotetion_book/internal/service"
	"quotetion_book/mocks"
	"strings"
	"testing"
)

func TestQuotesGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		postRequest    string
		mockSetup      func(m *mocks.MockQuotationService)
		expectedStatus int
		expectedBody   string
	}{
		{
			postRequest: `{"author": "Leo Tolstoy","quote": "Everyone thinks of changing the world, but no one thinks of changing himself."}`,
			mockSetup: func(m *mocks.MockQuotationService) {
				m.EXPECT().AddQuote(models.QuoteBook{
					Author: "Leo Tolstoy",
					Quote:  "Everyone thinks of changing the world, but no one thinks of changing himself.",
				})
			},
			expectedStatus: http.StatusCreated,
		},
		{
			mockSetup:      func(m *mocks.MockQuotationService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"ошибка":"неправильный формат данных"}`,
		},
		{
			postRequest:    `{}`,
			mockSetup:      func(m *mocks.MockQuotationService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"ошибка":"неправильный формат данных"}`,
		},
	}

	for _, tc := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockQuotationService := mocks.NewMockQuotationService(ctrl)
		tc.mockSetup(mockQuotationService)

		mockService := &service.Service{QuotationService: mockQuotationService}
		h := handlers.NewHandlers(mockService)

		router := gin.Default()
		router.POST("/quotes", h.QuotesPost)

		req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(tc.postRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedStatus, w.Code)
		assert.Contains(t, w.Body.String(), tc.expectedBody)
	}
}
