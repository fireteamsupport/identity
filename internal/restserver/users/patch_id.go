package userroutes

import (
    "github.com/labstack/echo/v4"
)

func (u *user) PatchId(c echo.Context) error {
    return c.String(200, "")
}
