package middleware

import (
    "github.com/fireteamsupport/identity/internal/restserver/structs"
    "github.com/labstack/echo/v4"
    "net/http"
)

func (m *Middleware) AuthZ(level int) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            user := c.Get("user").(*structs.User)

            log.Infof("Running AuthZ for user: %d", user.UID)

            if level > user.Role {
                return c.JSON(http.StatusUnauthorized, &structs.Message{
                    Code: http.StatusUnauthorized,
                    Message: "Sorry you are not allowed to use this endpoint",
                })
            }

            if user.Banned {
                return c.JSON(http.StatusUnauthorized, &structs.Message{
                    Code: http.StatusUnauthorized,
                    Message: "Sorry but it looks like you are banned",
                })
            }

            log.Infof("Looks like user %d is allowed to access this endpoint", user.UID)

            return next(c)
        }
    }
}

func (m *Middleware) AuthZDefault(next echo.HandlerFunc) echo.HandlerFunc {
    return m.AuthZ(0)(next)
}
