package authroutes

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/utils"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
)

func (a *auth) Logout(c echo.Context) error {
    err, token := utils.BearerExtractor(c)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusUnauthorized, &structs.Message{
            Message: "Invalid token",
        })
    }


    if err := a.RT.Delete(token); err != nil {
        log.Error(err)
        return c.String(403, "Invalid refresh token")
    }

    return c.JSON(http.StatusOK, &structs.Message{
        Message: "Ok Goodbye :)",
    })
}
