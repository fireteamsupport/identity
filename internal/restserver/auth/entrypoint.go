package authserver

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
    "github.com/fireteamsupport/identity/internal/restserver/middleware"
)

var log = logging.New()

type auth struct {
    *restutils.Options
}

func New(g *echo.Group, opts *restutils.Options) error {
    a := &auth{opts}

    g.POST("/login", a.Login)
    g.POST("/logout", a.Logout, middleware.UserAuth)
    g.POST("/register", a.Register)
    g.POST("/refresh", a.RefreshToken, middleware.UserAuth)
    g.POST("/passwordrest", a.PasswordReset)
    g.POST("/recover", a.Recover)
}
