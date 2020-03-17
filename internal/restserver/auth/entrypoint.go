package authserver

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
    "github.com/fireteamsupport/identity/internal/utils"

)

var (
    log = logging.New()
    v = utils.Validator()
)

type auth struct {
    *restutils.Options
}

func New(g *echo.Group, opts *restutils.Options) error {
    a := &auth{opts}

    g.POST("/login", a.Login)
    g.POST("/logout", a.Logout)
    g.POST("/register", a.Register)
    g.POST("/refresh", a.RefreshToken)
    //g.POST("/passwordrest", a.PasswordReset)
    //g.POST("/recover", a.Recover)
    return nil
}
