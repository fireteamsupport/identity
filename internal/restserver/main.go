package restserver

import (
    "github.com/arturoguerra/go-logging"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/profiles/internal/natsclient"
    middleware "github.com/fireteamsupport/profiles/internal/restserver/middleware"
    auth  "github.com/fireteamsupport/profiles/internal/restserver/auth"
    users "github.com/fireteamsupport/profiles/internal/restserver/users"
)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func New(e *echo.Echo, n nats.Client) (*echo.Echo, error) {
    e := echo.New()
    base := e.Group(baseURI)
    midlware := middleware.New()

    authgrp := base.Group("/auth")
    usrgrp := base.Group("/users", midlware.UserAuth)

    auth.New(authgrp, n)
    users.New(usrgrp, n)

    return e, nil
}
