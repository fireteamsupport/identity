package database

import (
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) NewAccountVerification(uid int64) *models.AccountVerification {
    log.Infof("Creating new account verification token for: %d", uid)

    av := &models.AccountVerification{
        UID: uid,
    }

    c.Create(av)

    return av
}
