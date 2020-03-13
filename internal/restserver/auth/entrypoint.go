package authserver

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/profiles/internal/restserver/structs"
    "github.com/fireteamsupport/profiles/internal/natsclient"
    "github.com/fireteamsupport/profiles/internal/database"
    "github.com/fireteamsupport/profiles/internal/jwtmanager"
    "github.com/fireteamsupport/profiles/internal/rtmanager"
    "github.com/fireteamsupport/profiles/internal/utils"
)

var log = logging.New()

type auth struct {
    DBClient database.Client
    NATSClient natsclient.Client
    JWTManager jwtmanager.JWTManager
    RTManager  rtmanager.RTManager
}

func New(g *echo.Group, opts *utils.Options) error {

    a := &auth{
        DBClient: opts.DBClient,
        NATSClient: opts.NATSClient,
        JWTManager: opts.JWTManager,
        RTManager: opts.RTManager,
    }

    g.POST("/login", a.Login)
    g.POST("/logout", a.Logout, mdlware.UserAuth)
    g.POST("/register", a.Register)
    g.POST("/refresh", a.Refresh, mdlware.UserAuth)
    g.POST("/passwordrest", a.PasswordReset)
    g.POST("/recover", a.Recover)
}
