package authroutes

import (
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/fireteamsupport/identity/internal/structs"
    "github.com/fireteamsupport/identity/internal/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

func (a *auth) RefreshToken(c echo.Context) error {
    err, header := utils.BearerExtractor(c)

    err, token := a.RT.Get(header)
    if err != nil {
        log.Error(err)
        return c.String(404, "Invalid refresh token")
    }

    err, dbuser := a.Store.User.GetId(token.UID)
    if err != nil {
        log.Error(err)
        return c.String(404, "Missing user")
    }

    err = a.RT.Delete(token.Token)
    if err != nil {
        log.Error(err)
        return c.String(500, "Error deleting expired refresh token")
    }

    user := &jwtmanager.User{
        UID: dbuser.UID,
        Email: dbuser.Email,
        Username: dbuser.Username,
    }

    atoken, err := a.JWT.Sign(user)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating user token try again later")
    }

    err, rtoken := a.RT.Create(user.UID, c.RealIP())
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating refresh token")
    }

    return c.JSON(http.StatusOK, &structs.RespRefresh{
        AccessToken: atoken,
        RefreshToken: rtoken,
        TokenType: "Bearer",
    })
}
