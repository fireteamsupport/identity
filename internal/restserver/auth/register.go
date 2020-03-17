package authserver

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/structs"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
)

func (a *auth) Register(c echo.Context) error {
    u := new(structs.ReqRegister)
    if err := c.Bind(u); err != nil {
        return c.String(500, "LOL")
    }

    if err := v.Struct(u); err != nil {
        log.Error(err)
        return c.String(400, "Unable to parse your input")
    }

    err, dbuser := a.DB.RegisterUser(u.Username, u.Email, u.Password)
    if err != nil {
        log.Error(err)
        return c.JSON(403, map[string]string{
            "code": "403",
            "message": "user exists",
        })
    }

    user := &jwtmanager.User{
        UID: dbuser.UID,
        Email: dbuser.Email,
        Username: dbuser.Username,
    }

    token, err := a.JWTMgmt.Sign(user)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating user token")
    }

    err, refreshtoken := a.RTMgmt.Create(user.UID, "4.2.4.2")
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating refresh token")
    }

    return c.JSON(http.StatusOK, &structs.RespRegister{
        AccessToken: token,
        RefreshToken: refreshtoken,
        TokenType: "Bearer",
    })
}
