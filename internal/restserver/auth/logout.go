package authroutes

import (
    "strings"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
)

func (a *auth) Logout(c echo.Context) error {
    value := c.Request().Header.Get("Authorization")
    schema := "Bearer"
    l := len(schema)
    if l+1 >= len(value) && value[:l] != schema {
        return c.JSON(http.StatusUnauthorized, &structs.Message{
            Message: "Missing or invalid token",
        })
    }

    token := strings.TrimSpace(value[l+1:])


    if err := a.RTMgmt.Delete(token); err != nil {
        log.Error(err)
        return c.String(403, "Invalid refresh token")
    }

    return c.JSON(http.StatusOK, &structs.Message{
        Message: "Ok Goodbye :)",
    })
}
