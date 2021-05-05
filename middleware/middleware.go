package middleware

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/saaitt/marketingManagerPod/handler"
	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/service"
)

func BasicAuthMiddlewareForUserType(userService service.UserService, expectedUserType string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		user, err := userService.Authenticate(request.AuthenticationRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			return false, err
		}
		if user.UserType != expectedUserType {
			return false, fmt.Errorf("only users of type %s can use this endpoint", expectedUserType)
		}
		c.Set(handler.UserIdKey, user.ID)
		c.Set(handler.UserTypeKey, user.UserType)
		return true, nil
	})
}
