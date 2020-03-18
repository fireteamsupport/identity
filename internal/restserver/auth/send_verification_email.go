package authroutes

import (
    "fmt"
)

func generateEmailBody(url string) string {
    base := `
    <body>
      <a href="%s">Verify</a>
    </body>
    `

    return fmt.Sprintf(base, url)
}

func (a *auth) SendVerificationEmail(email, subject, code string) {
    baseURL := fmt.Sprintf("%s?code=%s", email, code)
    a.Email.Send(email, subject, generateEmailBody(baseURL))
}
