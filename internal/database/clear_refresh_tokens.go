package database

import "github.com/fireteamsupport/identity/internal/models"

func (c *client) ClearRefreshTokens(uid int64) (error) {
    log.Infof("Clearing all Refresh tokens for: %d", uid)
    c.Where("UID = ?", uid).Delete(&models.RefreshToken{})
    return nil
}
