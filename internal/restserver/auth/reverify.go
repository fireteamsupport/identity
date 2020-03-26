package authroutes

import (
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/structs"
    "net/http"
    "fmt"
)

func (a *auth) Reverify(c echo.Context) error {
    payload := new(structs.ReqReverify)
    if err := c.Bind(payload); err != nil {
        return c.String(400, "Invalid payload")
    }

    if err := a.Validate.Struct(payload); err != nil {
        return c.String(400, "Invalid payload")
    }

    log.Info(payload.Email)
    err, user := a.Store.User.GetEmail(payload.Email)
    if err != nil {
        return c.String(404, "User not found")
    }

    if user.Verified {
        return c.JSON(403, map[string]string{
            "message": "already verified",
        })
    }

    verify := a.Store.AccountVerification.New(user.UID)
    subject := fmt.Sprintf("Here's your verification email %s", user.Username)

    a.SendVerificationEmail(user.Email, subject, verify.Token)

    return c.JSON(http.StatusOK, map[string]string{
        "message": "Sent verification email",
    })
}
