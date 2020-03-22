package userroutes

import (
    "strconv"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/resetserver/structs"
)

func (u *user) PatchId(c echo.Context) error {
    id := c.Param("id")

    uid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusNotFound, &structs.Message{
            Code: http.StatusNotFound,
            Message: err.Error(),
        })
    }

    err, dbuser := u.DB.GetUser(uid)
    if err != nil {
        log.Error(err)
        return c.JSON(http.StatusNotFound, &structs.Message{
            Code: http.StatusNotFound,
            Message: "user not found",
        })
    }

    payload := new(structs.UsersReqPatchId)
    if err := c.Bind(payload); err != nil {
        return err
    }

    if err := v.Struct(payload); err != nil {
        return c.JSON(http.StatusBadRequest, &structs.Message{
            Code: http.StatusBadRequest,
            Message: err.Error(),
        })
    }

    dbuser.Role = payload.Role

    u.DB.Save(dbuser)

    return c.JSON(http.StatusOK, &structs.Message{
        Code: http.StatusOK,
        Message: "Updated user",
    })
}
