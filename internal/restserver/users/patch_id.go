package userroutes

import (
    "strconv"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/restserver/structs"
    "net/http"
)

type req_PatchId struct {
    Role int `json:"role"`
}

func (u *user) PatchId(c echo.Context) error {
    id := c.Param("id")

    uid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusNotFound, &structs.Message{
            Message: err.Error(),
        })
    }

    err, dbuser := u.Store.User.GetId(uid)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusNotFound, &structs.Message{
            Message: "user not found",
        })
    }

    payload := new(req_PatchId)
    if err := c.Bind(payload); err != nil {
        return err
    }

    if err := u.Validate.Struct(payload); err != nil {
        return c.JSON(http.StatusBadRequest, &structs.Message{
            Message: err.Error(),
        })
    }

    dbuser.Role = payload.Role

    u.Store.DB.Save(dbuser)

    return c.JSON(http.StatusOK, &structs.Message{
        Message: "Updated user",
    })
}
