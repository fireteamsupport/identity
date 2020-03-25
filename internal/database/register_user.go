package database

import (
    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) RegisterUser(username, email, password string) (error, *models.User) {
    u := models.User{}

    if c.Where("email = ?", email).First(&u).RecordNotFound() {
        newUser := &models.User{
            Username: username,
            Email: email,
            Password: password,
        }

        c.DB.Create(newUser)

        return nil, newUser
    }

    err := errors.New(errors.Exists, email)

    return err, nil
}
