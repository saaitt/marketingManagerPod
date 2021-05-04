package handler

import (
	"net/http"

	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/service"
	"strconv"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	Service service.ProductService
}
type MarketingHandler struct {
	Service service.MarketingService
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

func (i MarketingHandler) ListAllMarketingProducts(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrInternalServerError
	}
	resp, err := i.Service.ListAllMarketingProductsByUserID(id)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (i ProductHandler) CreateProduct(c echo.Context) error {
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

func (i MarketingHandler) Redirect(c echo.Context) error {
	uuid := c.Param("marketing_product")
	pageLink, err := i.Service.ResolvePage(uuid)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.Redirect(http.StatusTemporaryRedirect, pageLink)
}
