package authroutes

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/initializer"
)

var (
    log = logging.New()
)

type auth struct {
    *initializer.Rest
}

func New(g *echo.Group, opts *initializer.Rest) error {
    a := &auth{opts}

    g.POST("/login", a.Login)
    g.POST("/logout", a.Logout)
    g.POST("/register", a.Register)
    g.GET("/verify", a.Verify)
    g.POST("/refresh", a.RefreshToken)
    g.POST("/reverify", a.Reverify)
    g.POST("/passwordrest", a.PasswordReset)
    g.POST("/recoveraccount", a.RecoverAccount)
    return nil
}
