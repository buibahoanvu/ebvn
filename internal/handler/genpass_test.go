package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/buibahoanvu/ebvn/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testErr = errors.New("test error")

func TestGenPassHandler_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest     func(ctx *gin.Context)
		setupMockService func(ctx context.Context) *mocks.GenPass

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "success",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/genpass", nil)
			},

			setupMockService: func(ctx context.Context) *mocks.GenPass {
				serviceMock := mocks.NewGenPass(t)
				serviceMock.On("GeneratePassword", passwordlength).Return("123456789012", nil)
				return serviceMock
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"password":"123456789012"}`,
		},

		{
			name: "service failed",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/genpass", nil)
			},

			setupMockService: func(ctx context.Context) *mocks.GenPass {
				serviceMock := mocks.NewGenPass(t)
				serviceMock.On("GeneratePassword", passwordlength).Return("", testErr)
				return serviceMock
			},

			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: `{"error":"Internal Server Err"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			tc.setupRequest(ctx)

			mockSvc := tc.setupMockService(ctx)
			testHandler := NewGenPass(mockSvc)

			testHandler.GeneratePassword(ctx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
