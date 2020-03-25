package restserver

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"

    auth  "github.com/fireteamsupport/identity/internal/restserver/auth"
    users  "github.com/fireteamsupport/identity/internal/restserver/users"
    restutils  "github.com/fireteamsupport/identity/internal/restserver/utils"
    middleware "github.com/fireteamsupport/identity/internal/restserver/middleware"
)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func NewDefault(opts *restutils.Options) (error, *echo.Echo) {
    err, cfg := NewEnvConfig()
    if err != nil {
        return err, nil
    }

    err, e := New(cfg, opts)
    if err != nil {
        return err, nil
    }

    return nil, e
}


func New(cfg *Config, opts *restutils.Options) (error, *echo.Echo) {
    e := echo.New()
    baseapi := e.Group(baseURI)

    m := middleware.New(opts.JWTMgmt)

    authgrp := baseapi.Group("/auth")
    auth.New(authgrp, opts)

    ugrp := baseapi.Group("/users", m.AuthN)
    users.New(ugrp, opts)


    e.GET("/healthz", func(c echo.Context) error {
        return c.String(200, "Everything is okay thx :)")
    })

    go func() {
        if err := e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
            e.Logger.Info(err.Error())
        }
    }()

    return nil, e
}
