package http

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var(
	e *echo.Echo
)

func TestNewEcho(t *testing.T) {
	// userRepository나 commentUC를 사용할 일이 없으므로 nil을 전달해도 된다.
	e = NewEcho(nil, nil)
	t.Run("Find /api/comments route handler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/comments", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)
		assert.NotEqual(t, http.StatusNotFound, rec.Code)

		req = httptest.NewRequest(http.MethodGet, "/api/comments", nil)
		rec = httptest.NewRecorder()
		_ = e.NewContext(req, rec)
		assert.NotEqual(t, http.StatusNotFound, rec.Code)

		req = httptest.NewRequest(http.MethodGet, "/api/comments/1", nil)
		rec = httptest.NewRecorder()
		_ = e.NewContext(req, rec)
		assert.NotEqual(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Health check for /healthz", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
