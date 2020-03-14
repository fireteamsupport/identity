package authserver

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/events"
    "github.com/fireteamsupport/identity/internal/database"
)

var log = logging.New()

type auth struct {
    DB database.Client
    Events events.Channels
}

func New(g *echo.Group, db database.Client, events events.Channels) error {

    a := &auth{
        DB: db,
        Events: events,
    }

    g.POST("/login", a.Login)
    g.POST("/logout", a.Logout, mdlware.UserAuth)
    g.POST("/register", a.Register)
    g.POST("/refresh", a.Refresh, mdlware.UserAuth)
    g.POST("/passwordrest", a.PasswordReset)
    g.POST("/recover", a.Recover)
}
