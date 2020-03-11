package restusers

import (
    "github.com/fireteamsupport/internal/natsclient"
    "github.com/fireteamsupport/internal/database"
    "github.com/labstack/echo/v4"
)

type (
    users struct {
       NClient: nats.Client
       DBClient: database.Client
    }
)

func New(g *echo.Group, dbclient database.Client, nclient nats.Client) {
    u := &users{
        NClient: nclient,
        DBClient: dbclient,
    }

    g.GET("/me", u.GetMe)
    g.PATCH("/me", u.PatchMe)
}
