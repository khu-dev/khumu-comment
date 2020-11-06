// 아직 뭘 넣어야할지 모르겠다.
package http

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var(
	e *echo.Echo
)

func TestNewEcho(t *testing.T) {
}

func handlerFuncMock(c echo.Context) error{
	return c.String(200, "found route")
}

func generatePathExistsTest(t *testing.T, target, method string, body io.Reader) func(t * testing.T){
	return func(t *testing.T){
		req := httptest.NewRequest(method, target, body)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		_ = handlerFuncMock(context)
		assert.NotEqual(t, http.StatusNotFound, rec.Code)
	}
}

func generatePathNotExistsTest(t *testing.T, target, method string, body io.Reader) func(t * testing.T){
	return func(t *testing.T){
		req := httptest.NewRequest(method, target, body)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		_ = handlerFuncMock(context)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}