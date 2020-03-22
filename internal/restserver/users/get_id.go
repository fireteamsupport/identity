package userroutes

import (
    "github.com/labstack/echo/v4"
)

func (u *user) GetId(c echo.Context) error {
    id := c.Param("id")
    uid, err = strconv.ParseInt(id, 10, 64)
    if err != nil {
        log.Error(err)
        return c.String(404, "Invalid ID")
    }

    err, dbuser := u.DB.GetUser(uid)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusNotFound, &structs.Message{
            Code: http.StatusNotFound,
            message: "user not found",
        })
    }

    return c.JSON(http.StatusOK, &structs.UsersRespGetId{
        UID: dbuser.UID,
        Username: dbuser.Username,
        Role: dbuser.Role,
        Banned: dbuser.Banned,
        Verified: dbuser.Verified,
    })
}
