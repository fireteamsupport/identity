package databas

import (
    "github.com/fireteamsupport/profiles/internal/errors"
)

func (c *client) RegisterUser(username, email, password string) (error, *User) {
    u := User{}

    if c.Where("email = ?").First(&u).RecordNotFound() {
        newUser := &User{
            Username: username,
            Email: email,
            Password: password,
        }

        c.DB.Create(newUser)

        return nil, newUser
    }

    err := errors.AlreadyExists(email)

    return err, nil
}
