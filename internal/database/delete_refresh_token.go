package database

import "github.com/fireteamsupport/identity/internal/errors"

func (c *client) DeleteRefreshToken(token string) (error) {
    rt := RefreshToken{}
    log.Infof("Deleting Refresh token: %d", uid)
    db.Where("Token = ?", token).Delete(&RefreshToken{})
    return nil
}
