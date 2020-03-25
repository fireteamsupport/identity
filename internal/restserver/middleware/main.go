package middleware

import (
    "fmt"
    "strings"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/arturoguerra/go-logging"
)

var log = logging.New()

type (
    Extractor func(echo.Context) (error, string)

    Middleware struct {
        JWTMgmt jwtmanager.JWTManager
    }
)

func New(jwt jwtmanager.JWTManager) *Middleware {
    return &Middleware{jwt}
}

func valueFromHeader(header, schema string) Extractor {
    return func(c echo.Context) (error, string) {
        value := c.Request().Header.Get(header)
        l := len(schema)
        if len(value) >= l+1 && value[:l] == schema {
            return nil, strings.TrimSpace(value[l+1:])
        }

        return fmt.Errorf("%s header not found", schema), ""
    }
}


