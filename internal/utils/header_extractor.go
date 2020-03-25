package utils

import (
    "fmt"
    "strings"
    "github.com/labstack/echo/v4"
)

var (
    BearerExtractor = ValueFromHeader("Authorization", "Bearer")
)

type Extractor func(echo.Context) (error, string)

func ValueFromHeader(header, schema string) Extractor {
    return func(c echo.Context) (error, string) {
        value := c.Request().Header.Get(header)
        l := len(schema)
        if len(value) >= l+1 && value[:l] == schema {
            return nil, strings.TrimSpace(value[l+1:])
        }

        return fmt.Errorf("%s header not found", schema), ""
    }
}
