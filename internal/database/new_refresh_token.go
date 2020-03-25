package database

import (
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) NewRefreshToken(uid int64, ip string) *models.RefreshToken {
    log.Infof("Creating new refresh token for: %d", uid)

    rt := &models.RefreshToken{
        UID: uid,
        IP: ip,
    }

    c.Create(rt)

    return rt
}
