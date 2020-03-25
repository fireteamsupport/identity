package authroutes

import (
    "github.com/labstack/echo/v4"
    "github.com/fireteamsupport/identity/internal/structs"
    "net/http"
    "fmt"
)

const passwordResetTemplate = `
<body>
<a href="%s">Reset Password</a>
</body>
`

func (a *auth) PassResetEmail(email, username, code string) {
    subject := fmt.Sprintf("Here's the password rest you requested %s", username)

    body := fmt.Sprintf(passwordResetTemplate, code)

    a.Email.Send(email, subject, body)
}

func (a *auth) PasswordReset(c echo.Context) error {
    payload := new(structs.ReqPasswordReset)

    if err := c.Bind(payload); err != nil {
        log.Error(err)
        return c.String(400, "Invalid payload")
    }

    if err := a.Validate.Struct(payload); err != nil {
        log.Error(err)
        return c.String(400, "Invaliad payload")
    }

    log.Info(payload)
    err, user := a.DB.UserLogin(payload.Email)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusOK, "may or may not exists we will never know")
    }

    reset := a.DB.NewPasswordReset(user.UID)

    a.PassResetEmail(user.Email, user.Username, reset.Token)

    return c.String(http.StatusOK, "may or may not exists we will never know")
}
