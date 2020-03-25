package database

import (
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) GetRefreshTokens(uid int64) (error, []*models.RefreshToken) {
    rts := make([]*models.RefreshToken, 0)
    log.Infof("Getting Refresh tokens for: %d", uid)
    c.Where("UID = ?", uid).First(&rts)
    return nil, rts
}
