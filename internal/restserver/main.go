package restserver

import (
    "github.com/arturoguerra/go-logging"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/events"

    middleware "github.com/fireteamsupport/identity/internal/restserver/middleware"
    auth  "github.com/fireteamsupport/identity/internal/restserver/auth"
    //users "github.com/fireteamsupport/profiles/internal/restserver/users"
)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func New(e *echo.Echo, db database.Client, events events.Client) (*echo.Echo, error) {
    e := echo.New()
    base := e.Group(baseURI)
    midlware := middleware.New()

    authgrp := base.Group("/auth")
    auth.New(authgrp, db, events)

    //usrgrp := base.Group("/users", midlware.UserAuth)
    //users.New(usrgrp, opts)

    return e, nil
}
