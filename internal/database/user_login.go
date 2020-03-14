package database

/*
Gets a single user from our database
*/

import (
    "github.com/fireteamsupport/identity/internal/errors"
)

func (c *client) UserLogin(email string) (error, *User) {
    u := User{}
    log.Infof("Getting user for Email: %s", email)
    if c.Where("Email = ?", email).First(&u).RecordNotFound() {
        return errors.New(errors.NotFound, email), nil
    }

    return nil, &u
}
