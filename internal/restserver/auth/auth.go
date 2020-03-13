package restserver

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

func NewAuthModule(g *echo.Group, opts *utils.Options) error {

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

func (a *auth) Login(c echo.Context) error {
    payload := new(structs.ReqLogin)
    if err := c.Bind(payload); err != nil {
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    if err = a.Validator.Struct(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Error validating")
    }

    err, dbuser := a.DBClient.UserLogin(payload.Email)
    if err != nil {
        if err == errors.NotFound {
            return c.String(http.StatusNotFound, "User not found")
        }

        log.Error(err)
        return c.String(http.StatusInternalServerError, "Unable to fetch user try again later")
    }


    if dbuser.Password != payload.Password {
        return c.String(401, "Invalid email and/or password")
    }

    user := &a.JWTManager.User{
        UID: dbuser.UID,
        Email: dbuser.Email,
        Username: dbuser.Username,
    }

    token, err := a.JWTManager.Sign(user)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating user token try again later")
    }

    refreshtoken, err := a.RTManager.New(user.UID, "")
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating refresh token")
    }

    return c.JSON(http.StatusOK, &structs.RespLogin{
        AccessToken: token,
        RefreshToken: refreshtoken,
    })
}

func (a *auth) Logout(c echo.Context) error {
}

func (a *auth) Register(c echo.Context) error {
}

func (a *auth) PasswordReset(c echo.Context) error {
}

func (a *auth) Recover(c echo.Context) error {
}
