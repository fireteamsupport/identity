package database

import (
    "time"
    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
)

func (c *client) GetAccountVerification(token string) (error, *models.AccountVerification) {
    av := models.AccountVerification{}
    log.Infof("Getting Account verification token: %s", token)
    if c.Where("Token = ?", token).First(&av).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= av.ExpiresAt.Unix() {
        return errors.New(errors.Expired, av.Token), nil
    }

    return nil, &av
}
