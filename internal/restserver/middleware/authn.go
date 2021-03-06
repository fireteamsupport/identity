package middleware

import (
    "github.com/fireteamsupport/identity/internal/restserver/structs"
    "github.com/fireteamsupport/identity/internal/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

func (m *Middleware) AuthN(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        err, token := utils.BearerExtractor(c)
        if err != nil {
            log.Error(err)
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "message": "missing or invalid token",
            })
        }

        err, claims := m.JWTMgmt.Decrypt(token)
        if err != nil {
            log.Error(err)
            return c.JSON(403, map[string]string{
                "message": "Invalid token",
            })
        }

        user := &structs.User{
            UID: claims.UID,
            Username: claims.Username,
            Email: claims.Email,
            Role: claims.Role,
            Banned: claims.Banned,
        }

        c.Set("user", user)

        return next(c)
    }
}
