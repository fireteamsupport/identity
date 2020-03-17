package restserver

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/config"

    auth  "github.com/fireteamsupport/identity/internal/restserver/auth"
    restutils  "github.com/fireteamsupport/identity/internal/restserver/utils"
)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func New(cfg *config.EchoConfig, opts *restutils.Options) (error, *echo.Echo) {
    e := echo.New()
    baseapi := e.Group(baseURI)

    authgrp := baseapi.Group("/auth")
    auth.New(authgrp, opts)

    go func() {
        if err := e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
            e.Logger.Info(err.Error())
        }
    }()

    return nil, e
}
