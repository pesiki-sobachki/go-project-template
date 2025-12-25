package v1_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/shanth1/gotools/log"
	v1 "github.com/shanth1/template/internal/adapter/driving/http/v1"
	"github.com/shanth1/template/internal/core/port/mocks"
)

func TestHandler_HealthCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)

	logger := log.New(log.WithLevel(log.LevelDisabled))

	h := v1.NewHandler(mockService, logger)

	t.Run("Success", func(t *testing.T) {
		mockService.EXPECT().HealthCheck(gomock.Any()).Return(nil)

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()

		h.HealthCheck(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Failure", func(t *testing.T) {
		mockService.EXPECT().HealthCheck(gomock.Any()).Return(errors.New("db down"))

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()

		h.HealthCheck(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})
}
