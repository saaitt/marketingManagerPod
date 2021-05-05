package handler

import (
	"net/http"

	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/service"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	Service service.ProductAdminService
}
type MarketingHandler struct {
	Service service.MarketingService
}

const (
	UserTypeKey = "userTypeKey"
	UserIdKey   = "userId"
)

func getUserId(c echo.Context) int {
	return c.Get(UserIdKey).(int)
}
func (i ProductHandler) Create(c echo.Context) error {
	req := request.CreateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	req.UserID = getUserId(c)
	resp, err := i.Service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (i ProductHandler) FindByUser(c echo.Context) error {
	resp, err := i.Service.FindByUser(getUserId(c))
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (i MarketingHandler) FindByUser(c echo.Context) error {
	resp, err := i.Service.FindByUserID(getUserId(c))
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}

func (i MarketingHandler) CreateProduct(c echo.Context) error {
	req := request.CreateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp, err := i.Service.CreateProduct(req.ID, getUserId(c))
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

type UserHandler struct {
	Service service.UserService
}

func (i UserHandler) Create(c echo.Context) error {
	req := request.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp, err := i.Service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, resp)
}
