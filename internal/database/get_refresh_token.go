package database

import "github.com/fireteamsupport/identity/internal/errors"

func (c *client) GetRefreshToken(token string) (error, *RefreshToken) {
    rt := RefreshToken{}
    log.Infof("Getting Refresh token: %s", token)
    if c.WHere("token = ?").First(&rt).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    return nil, &rt
}
