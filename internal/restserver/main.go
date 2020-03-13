package restserver

import (
    "github.com/arturoguerra/go-logging"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/utils"
    "github.com/fireteamsupport/identity/internal/natsclient"
    "github.com/fireteamsupport/identity/internal/database"
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

func New(e *echo.Echo, opts *utils.Options) (*echo.Echo, error) {
    e := echo.New()
    base := e.Group(baseURI)
    midlware := middleware.New()

    authgrp := base.Group("/auth")
    auth.New(authgrp, opts)

    //usrgrp := base.Group("/users", midlware.UserAuth)
    //users.New(usrgrp, opts)

    return e, nil
}
