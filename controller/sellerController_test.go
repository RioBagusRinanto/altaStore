package controller

import (
	"altaStore/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetSellersController(t *testing.T) {
	// var testCases = []struct {
	// 	name                string
	// 	path                string
	// 	expectStatus        int
	// 	expectBodyStartWith string
	// }{
	// 	{
	// 		name:                "sukses get seller",
	// 		path:                "/sellers",
	// 		expectStatus:        http.StatusOK,
	// 		expectBodyStartWith: "{\"status\":\"success\"",
	// 	},
	// 	{
	// 		name:                "gagal get seller",
	// 		path:                "/sellers",
	// 		expectStatus:        http.StatusBadRequest,
	// 		expectBodyStartWith: "",
	// 	},
	// }

	// e := InitEcho()
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)

	// for _, tesCase := range testCases {
	// 	c.SetPath(tesCase.path)

	// 	//assertion
	// 	if assert.NoError(t, GetSellersController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 		body := rec.Body.String()
	// 		assert.True(t, strings.HasPrefix(body, tesCase.expectBodyStartWith))
	// 	}
	// }

	t.Run("test case 1, sukses get seller", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/sellers")
		if assert.NoError(t, GetSellersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("test case 2, gagal get seller", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/sellers")

		if assert.NoError(t, GetSellersController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestInsertSellersController(t *testing.T) {

	t.Run("test case 1, sukses insert seller", func(t *testing.T) {
		entryseller := `{"username":"seller2", "password":"seller2", "alamat":"alamat seller2", "nama_toko":"toko2", "email":"email2@email.com"}`
		e := InitEcho()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(entryseller))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sellers")

		if assert.NoError(t, InsertSellerController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("test case 2, gagal insert seller", func(t *testing.T) {
		entryseller := `{"id_seller":1,"username":"seller2", "password":"seller2", "alamat":"alamat seller2", "nama_toko":"toko2", "email":"email2@email.com"}`
		e := InitEcho()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(entryseller))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sellers")

		if assert.NoError(t, InsertSellerController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

}
