package userroutes

import (
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
    "net/http"
)

func (u *user) GetME(c echo.Context) error {
    user := c.Get("user").(*structs.User)

    err, dbuser := u.Store.User.GetId(user.UID)
    if err != nil {
        log.Error(err)
        return err
    }

    return c.JSON(http.StatusOK, &resp_User{
        UID: dbuser.UID,
        Username: dbuser.Username,
        Email: dbuser.Email,
        Banned: dbuser.Banned,
        Verified: dbuser.Verified,
        Role: dbuser.Role,
    })
}
