package authroutes

import (
    "github.com/fireteamsupport/identity/internal/structs"
    "github.com/labstack/echo/v4"
    "net/http"
)

func (a *auth) RecoverAccount(c echo.Context) error {
    payload := new(structs.ReqRecoverAccount)
    if err := c.Bind(payload); err != nil {
        log.Error(err)
        return c.String(400, "Invalid payload")
    }

    if err := v.Struct(payload); err != nil {
        log.Error(err)
        return c.String(400, "Invalid payload")
    }

    err, entry := a.DB.GetPasswordReset(payload.Code)
    if err != nil {
        log.Error(err)
        return c.String(403, "Invalid code")
    }

    err, user := a.DB.GetUser(entry.UID)
    if err != nil {
        log.Error(err)
        return c.String(403, "Invalid user")
    }

    user.Password = payload.Password
    a.DB.Save(user)

    a.DB.Delete(entry)

    return c.JSON(http.StatusOK, map[string]string{
        "message": "password updated",
    })
}
