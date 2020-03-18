package authroutes

import (
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/fireteamsupport/identity/internal/structs"
    "github.com/labstack/echo/v4"
    "net/http"
)

func (a *auth) RefreshToken(c echo.Context) error {
    u := new(structs.ReqRefresh)
    if err := c.Bind(u); err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Invalid payload")
    }

    if err := v.Struct(u); err != nil {
        log.Error(err)
        return c.String(400, "Invalid payload")
    }

    err, token := a.RTMgmt.Get(u.Token)
    if err != nil {
        log.Error(err)
        return c.String(404, "Invalid refresh token")
    }

    err, dbuser := a.DB.GetUser(token.UID)
    if err != nil {
        log.Error(err)
        return c.String(404, "Missing user")
    }

    err = a.RTMgmt.Delete(token.Token)
    if err != nil {
        log.Error(err)
        return c.String(500, "Error deleting expired refresh token")
    }

    user := &jwtmanager.User{
        UID: dbuser.UID,
        Email: dbuser.Email,
        Username: dbuser.Username,
    }

    atoken, err := a.JWTMgmt.Sign(user)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating user token try again later")
    }

    err, rtoken := a.RTMgmt.Create(user.UID, "5.534.55.5")
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
