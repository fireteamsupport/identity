package database

/*
Gets a single user from our database
*/

import (
    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) GetUser(uid int64) (error, *models.User) {
    u := models.User{}
    log.Infof("Getting user for UID: %d", uid)
    if c.Where("UID = ?", uid).First(&u).RecordNotFound() {
        return errors.New(errors.NotFound, uid), nil
    }

    return nil, &u
}
