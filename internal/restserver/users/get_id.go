package userroutes

import (
    "github.com/labstack/echo/v4"
    "strconv"
    "net/http"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
)


func (u *user) GetId(c echo.Context) error {
    id := c.Param("id")
    uid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        log.Error(err)
        return c.String(404, "Invalid ID")
    }

    err, dbuser := u.Store.User.GetId(uid)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusNotFound, &structs.Message{
            Code: http.StatusNotFound,
            Message: "user not found",
        })
    }

    return c.JSON(http.StatusOK, &resp_User{
        UID: dbuser.UID,
        Username: dbuser.Username,
        Email: dbuser.Email,
        Role: dbuser.Role,
        Banned: dbuser.Banned,
        Verified: dbuser.Verified,
    })
}
