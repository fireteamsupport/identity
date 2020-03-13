package database

/*
Gets a single user from our database
*/

import (
    "github.com/fireteamsupport/identity/errors"
)

func (c *client) GetUser(uid int64) (error, *User) {
    u := User{}
    log.Infof("Getting user for UID: %d", uid)
    if c.Where("UID = ?", uid).First(&u).RecordNotFound() {
        return errors.New(errors.NotFound, uid)
    }

    return nil, &u
