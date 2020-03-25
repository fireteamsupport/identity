package authroutes

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
)

type (
    req_Login struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    resp_Login struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
        TokenType    string `json:"token_type"`
    }
)

func (a *auth) Login(c echo.Context) error {
    payload := new(req_Login)
    if err := c.Bind(payload); err != nil {
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    if err := a.Validate.Struct(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Error validating")
    }

    err, dbuser := a.DB.UserLogin(payload.Email)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusNotFound, "User not found")
    }

    if !dbuser.ValidPassword(payload.Password) {
        return c.String(401, "Invalid email and/or password")
    }

    if !dbuser.Verified {
        return c.JSON(403, map[string]string{
            "message": "Account is not verified",
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
        return c.String(http.StatusInternalServerError, "Error creating user token try again later")
    }

    err, refreshtoken := a.RTMgmt.Create(user.UID, c.RealIP())
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Error creating refresh token")
    }

    return c.JSON(http.StatusOK, &resp_Login{
        AccessToken: token,
        RefreshToken: refreshtoken,
        TokenType: "Bearer",
    })
}
