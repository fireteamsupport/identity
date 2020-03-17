package authserver

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/structs"
)

func (a *auth) Logout(c echo.Context) error {
    s := new(structs.ReqLogout)
    if err := c.Bind(s); err != nil {
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    if err := v.Struct(s); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    if err := a.RTMgmt.Delete(s.Token); err != nil {
        log.Error(err)
        return c.String(403, "Invalid refresh token")
    }

    return c.String(http.StatusOK, "Goodbye :)")
}
