package authserver

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/structs"
)

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
        if err.Code() == errors.NotFound {
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
