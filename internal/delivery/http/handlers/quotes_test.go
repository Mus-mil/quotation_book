package handlers_test

import (
	"errors"
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

func TestQuotesPost(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		postRequest    string
		mockSetup      func(m *mocks.MockQuotationService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "valid request",
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
			name:           "missing body",
			mockSetup:      func(m *mocks.MockQuotationService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"ошибка":"неправильный формат данных"}`,
		},
		{
			name:           "empty json body",
			postRequest:    `{}`,
			mockSetup:      func(m *mocks.MockQuotationService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"ошибка":"неправильный формат данных"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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
		})
	}
}

func TestQuotesGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		query          string
		mockSetup      func(m *mocks.MockQuotationService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:  "get all quotes",
			query: "",
			mockSetup: func(m *mocks.MockQuotationService) {
				m.EXPECT().GetAllQuotes().Return([]models.QuoteBookID{
					{ID: 1, Author: "Author", Quote: "Quote"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Author",
		},
		{
			name:  "get quotes by author",
			query: "?author=Leo",
			mockSetup: func(m *mocks.MockQuotationService) {
				m.EXPECT().GetQuotesFromAuthor("Leo").Return([]models.QuoteBookID{
					{ID: 2, Author: "Leo", Quote: "Quote"},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Leo",
		},
		{
			name:  "unknown author",
			query: "?author=Unknown",
			mockSetup: func(m *mocks.MockQuotationService) {
				m.EXPECT().GetQuotesFromAuthor("Unknown").Return(nil, nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "автора нет",
		},
		{
			name:  "db failure",
			query: "?author=Error",
			mockSetup: func(m *mocks.MockQuotationService) {
				m.EXPECT().GetQuotesFromAuthor("Error").Return(nil, errors.New("db fail"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "ошибка",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			MockQuotationService := mocks.NewMockQuotationService(ctrl)
			tc.mockSetup(MockQuotationService)

			svc := &service.Service{QuotationService: MockQuotationService}
			h := handlers.NewHandlers(svc)

			router := gin.Default()
			router.GET("/quotes", h.QuotesGet)

			req := httptest.NewRequest(http.MethodGet, "/quotes"+tc.query, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedBody)
		})
	}
}
