package authserver

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/fireteamsupport/identity/internal/structs"
)

func (a *auth) Login(c echo.Context) error {
    payload := new(structs.ReqLogin)
    if err := c.Bind(payload); err != nil {
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    if err := v.Struct(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Error validating")
    }

    err, dbuser := a.DB.UserLogin(payload.Email)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusNotFound, "User not found")
    }


    if dbuser.Password != payload.Password {
        return c.String(401, "Invalid email and/or password")
    }

    user := &jwtmanager.User{
        UID: dbuser.UID,
        Email: dbuser.Email,
        Username: dbuser.Username,
    }

    token, err := a.JWTMgmt.Sign(user)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating user token try again later")
    }

    refreshtoken, err := a.RTMgmt.Create(user.UID, "423.42.3.4")
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating refresh token")
    }

    return c.JSON(http.StatusOK, &structs.RespLogin{
        AccessToken: token,
        RefreshToken: refreshtoken,
    })
}
