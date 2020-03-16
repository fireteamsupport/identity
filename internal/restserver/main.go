package restserver

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"

    auth  "github.com/fireteamsupport/identity/internal/restserver/auth"
    restutils  "github.com/fireteamsupport/identity/internal/restserver/utils"
)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func New(e *echo.Echo, opts *restutils.Options) (*echo.Echo, error) {
    baseapi := e.Group(baseURI)

    authgrp := baseapi.Group("/auth")
    auth.New(authgrp, opts)

    return e, nil
}
