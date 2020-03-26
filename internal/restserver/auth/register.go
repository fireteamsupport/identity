package authroutes

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/structs"
)

func verifyEmailBody(code string) string {
    return "http://10.50.1.15:5000/api/v1/auth/verify?=" + code
}

func (a *auth) Register(c echo.Context) error {
    u := new(structs.ReqRegister)

    if err := c.Bind(u); err != nil {
        return c.String(500, "LOL")
    }

    if err := a.Validate.Struct(u); err != nil {
        log.Error(err)
        return c.String(400, "Unable to parse your input")
    }

    err, dbuser := a.Store.User.New(u.Username, u.Email, u.Password)
    if err != nil {
        log.Error(err)
        return c.JSON(403, map[string]string{
            "code": "403",
            "message": "user exists",
        })
    }

    verify := a.Store.AccountVerification.New(dbuser.UID)
    subject := fmt.Sprintf("Here's your verification email %s", dbuser.Username)

    a.SendVerificationEmail(dbuser.Email, subject, verify.Token)

    return c.JSON(http.StatusOK, &structs.RespRegister{
        Username: dbuser.Username,
        Email: dbuser.Email,
    })
}
