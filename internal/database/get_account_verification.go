package database

import (
    "time"
    "github.com/fireteamsupport/identity/internal/errors"
)

func (c *client) GetAccountVerification(token string) (error, *AccountVerification) {
    av := AccountVerification{}
    log.Infof("Getting Account verification token: %s", token)
    if c.Where("Token = ?", token).First(&av).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= av.ExpiresAt.Unix() {
        return errors.New(errors.Expired, av.Token), nil
    }

    return nil, &av
}
