package database

import (
    "time"
    "github.com/fireteamsupport/identity/internal/errors"
)

func (c *client) GetPasswordReset(token string) (error, *PasswordReset) {
    pr := PasswordReset{}
    log.Infof("Getting Password Reset: %s", token)
    if c.Where("Token = ?", token).First(&pr).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= pr.ExpiresAt.Unix() {
        return errors.New(errors.Expired, pr.Token), nil
    }

    return nil, &pr
}
