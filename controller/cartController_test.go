package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCartByIdCustController(t *testing.T) {

	t.Run("test case 1, sukses get cart", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/carts")
		c.SetParamNames("userid")
		c.SetParamValues("1")
		if assert.NoError(t, GetCartByIdCustController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("test case 2, gagal get cart", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/carts/:userid")
		c.SetParamNames("userid")
		c.SetParamValues("12")

		if assert.NoError(t, GetCartByIdCustController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestInsertCartController(t *testing.T) {

	t.Run("test case 1, sukses insert cart", func(t *testing.T) {
		entryseller := `{"total":200000, "status_pesanan":"cart", "id_customer":1}`
		e := InitEcho()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(entryseller))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/carts")

		if assert.NoError(t, InsertCartController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("test case 2, gagal insert carts", func(t *testing.T) {
		entryseller := `{"total":200000, "status_pesanan":"cart", "id_customer":12}`
		e := InitEcho()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(entryseller))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/carts")

		if assert.NoError(t, InsertCartController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

}
