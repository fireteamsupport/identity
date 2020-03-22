package userroutes

import (
    "github.com/labstack/echo/v4"
)

func (u *user) GetME(c echo.Context) error {
    user := c.Get("user").(*structs.User)

    err, dbuser := u.DB.GetUser(user.UID)
    if err != nil {
        log.Error(err)
        return err
    }

    return c.JSON(http.StatusOK, &structs.RespUserGetME{
        UID: dbuser.UID,
        Username: dbuser.Username,
        Email: dbuser.Email,
        Verified: dbuser.Verified,
        Role: dbuser.Role,
    })
}
