package database

import (
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) DeleteRefreshToken(token string) error {
    log.Infof("Deleting Refresh token: %s", token)
    c.Where("Token = ?", token).Delete(&models.RefreshToken{})
    return nil
}
