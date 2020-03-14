package restserver

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/events"

    auth  "github.com/fireteamsupport/identity/internal/restserver/auth"
)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func New(e *echo.Echo, opts *restutils.Options) (*echo.Echo, error) {
    e := echo.New()
    base := e.Group(baseURI)

    authgrp := base.Group("/auth")
    auth.New(authgrp, opts)

    return e, nil
}
