package database

import (
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) NewPasswordReset(uid int64) *models.PasswordReset {
    log.Infof("Creating new password reset for: %d", uid)

    ps := &models.PasswordReset{
        UID: uid,
    }

    c.Create(ps)

    return ps
}
