package userroutes

import (
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
)

func (u *user) PatchME(c echo.Context) error {
    user := c.Get("user").(*structs.User)

    err, dbuser := u.DB.GetUser(user.UID)
    if err != nil {
        log.Error(err)
        return err
    }

    payload := new(structs.ReqUsersPatchME)
    if err := c.Bind(payload); err != nil {
        log.Error(err)
        return err
    }

    if err := v.Struct(payload); err != nil {
        log.Error(err)
        return c.JSON(400, &structs.Message{
            Code: 400,
            Message: "invalid payload",
        })
    }

    if payload.OldPassword != dbuser.Password {
        return c.JSON(403, &structs.Message{
            Code: 403,
            Message: "invalid password",
        })
    }

    if payload.NewPassword != "" {
        dbuser.Password = payload.NewPassword
    }

    if payload.Username != "" {
        dbuser.Username = payload.Username
    }

    if payload.Email != "" {
        dbuser.Email = payload.Email
        dbuser.Verified = false
    }

    u.DB.Save(dbuser)

    return c.JSON(200, &structs.Message{
        Code: 200,
        Message: "updated user",
    })
}
