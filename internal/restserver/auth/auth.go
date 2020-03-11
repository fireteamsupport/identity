package restserver

import (
    "github.com/labstack/echo/v4"
)

type auth struct {}

func NewAuthModule(g *echo.Group) error {
    a := auth{}
    g.POST("/login", a.Login)
    g.POST("/logout", a.Logout, mdlware.UserAuth)
    g.POST("/register", a.Register)
    g.POST("/refresh", a.Refresh, mdlware.UserAuth)
    g.POST("/passwordrest", a.PasswordReset)
    g.POST("/recover", a.Recover)
}

func (a *auth) Login(c echo.Context) error {
}

func (a *auth) Logout(c echo.Context) error {
}

func (a *auth) Register(c echo.Context) error {
}

func (a *auth) PasswordReset(c echo.Context) error {
}

func (a *auth) Recover(c echo.Context) error {
}
