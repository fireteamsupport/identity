package authroutes

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func (a *auth) Verify(c echo.Context) error {
    value := c.QueryParam("code")
    err, entry := a.DB.GetAccountVerification(value)
    if err != nil {
        log.Error(err)
        return c.String(403, "Invalid code")
    }

    err, user := a.DB.GetUser(entry.UID)
    if err != nil {
        log.Error(err)
        return c.String(403, "Invalid user")
    }

    user.Verified = true
    a.DB.Save(user)

    return c.JSON(http.StatusOK, map[string]string{
        "message": "Account has been verified",
    })
}
