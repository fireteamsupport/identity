package restserver

import (
    "fmt"
    "github.com/labstack/echo/v4"
    basemiddleware "github.com/labstack/echo/v4/middleware"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/initializer"

    auth  "github.com/fireteamsupport/identity/internal/restserver/auth"
    users  "github.com/fireteamsupport/identity/internal/restserver/users"
    middleware "github.com/fireteamsupport/identity/internal/restserver/middleware"

)

const (
    baseURI = "/api/v1"
)

var (
    log = logging.New()
)

func NewDefault(opts *initializer.Rest) (*echo.Echo, error) {
    cfg, err := NewEnvConfig()
    if err != nil {
        return nil, err
    }

    e, err := New(cfg, opts)
    if err != nil {
        return nil, err
    }

    return e, nil
}


func New(cfg *Config, opts *initializer.Rest) (*echo.Echo, error) {
    e := echo.New()
    e.Use(basemiddleware.Logger())
    e.Use(basemiddleware.Recover())

    baseapi := e.Group(baseURI)

    m := middleware.New(opts.JWT)

    log.Info("Stuffs")
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

    return e, nil
}
