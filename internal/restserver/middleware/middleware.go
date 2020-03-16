package middleware

import (
    "github.com/fireteamsupport/identity/internal/structs"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
    "github.com/labstack/echo/v4"
    "strings"
)

type (
    middleware struct {
        *restutils.Options
    }

    Middleware interface {
        Auth(echo.HandlerFunc) echo.HandlerFunc
    }

    message struct {
        Code int `json:"code"`
        Message string `json:"message"`
    }
)

func New(opts *restutils.Options) Middleware {
    return &middleware{opts}
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        header := c.Request().Header.Get("Authorization")
        if header == "" {
            return c.JSON(401, &message{
                401,
                "Missing auth token",
            })
        }

        split := strings.Split(header, " ")
        if len(split) != 2 {
            return c.JSON(400, &message{
                400,
                "Malformed header",
            })
        }

        if split[0] != "Bearer" {
            return c.JSON(400, &message{
                400,
                "Malformed header",
            })
        }

        token := split[1]

        err, claims := m.JWTMgmt.Decrypt(token)
        if err != nil {
            return c.JSON(403, &message{
                403,
                "Invalid token",
            })
        }

        user := &structs.User{
            UID: claims.UID,
            Username: claims.Username,
            Email: claims.Email,
        }

        c.Set("user", user)
        return next(c)
    }
}
