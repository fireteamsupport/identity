package database

import "github.com/fireteamsupport/profiles/internal/errors"

func (c *client) ClearRefreshTokens(uid int64) (error) {
    rt := RefreshToken{}
    log.Infof("Clearing all Refresh tokens for: %d", uid)
    db.Where("UID = ?", uid).Delete(&RefreshToken{})
    return nil
}
