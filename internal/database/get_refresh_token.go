package database

import (
    "time"
    "github.com/fireteamsupport/identity/internal/errors"
)

func (c *client) GetRefreshToken(token string) (error, *RefreshToken) {
    rt := RefreshToken{}
    log.Infof("Getting Refresh token: %s", token)
    if c.Where("Token = ?", token).First(&rt).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= rt.ExpiresAt.Unix() {
        return errors.New(errors.Expired, rt.Token), nil
    }

    return nil, &rt
}
