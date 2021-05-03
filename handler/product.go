package handler

import (
	"net/http"

	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/service"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	Service service.ProductService
}

func (i ProductHandler) Create(c echo.Context) error {
	req := request.CreateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp, err := i.Service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (i ProductHandler) ListAll(c echo.Context) error {
	resp, err := i.Service.ListAll()
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}
