package userroutes

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
)

type (
    req_PatchME struct {
        Username    string `json:"username" validate:"required"`
        Email       string `json:"email"    validate:"required,email"`
        Password    string `json:"password" validate:"required"`
        NewPassword string `json:"new_password"`
    }
)

func (u *user) PatchME(c echo.Context) error {
    user := c.Get("user").(*structs.User)

    err, dbuser := u.Store.User.GetId(user.UID)
    if err != nil {
        log.Error(err)
        return err
    }

    payload := new(req_PatchME)
    if err := c.Bind(payload); err != nil {
        log.Error(err)
        return err
    }

    if err := u.Validate.Struct(payload); err != nil {
        log.Error(err)
        return c.JSON(400, &structs.Message{
            Code: 400,
            Message: "invalid payload",
        })
    }

    if payload.Password != dbuser.Password {
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
        if err, _ := u.Store.User.GetEmail(payload.Email); err == nil {
            return c.JSON(http.StatusUnauthorized, &structs.Message{
                Message: "Email already registered",
            })
        }

        dbuser.Email = payload.Email
        dbuser.Verified = false
    }

    u.Store.DB.Save(dbuser)

    return c.JSON(200, &structs.Message{
        Code: 200,
        Message: "updated user",
    })
}
