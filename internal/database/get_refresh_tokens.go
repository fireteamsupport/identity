package database

import "github.com/fireteamsupport/profiles/internal/errors"

func (c *client) GetRefreshTokens(uid int64) (error, []*RefreshToken) {
    rts := []*RefreshToken
    log.Infof("Getting Refresh token: %s", token)
    c.Where("token = ?").First(&rts)
    return nil, &rt
}
