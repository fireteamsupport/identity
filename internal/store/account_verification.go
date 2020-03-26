package store

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
    "github.com/fireteamsupport/identity/internal/database"
)

type (
    accountVerificationStore struct {
        *database.Client
    }

    AccountVerificationStore interface {
        New(int64) error
        GetByToken(string) (error, *models.AccountVerification)
        Save(interface{}) *gorm.DB
        Delete(interface{}, ...interface{}) *gorm.DB
    }
)

func NewAccountVerificationStore(db *database.Client) (*AccountVerificationStore, error) {
    db.AutoMigrate(&models.AccountVerification{})
    return &accountVerificationStore{db}, nil
}

func (store *accountVerificationStore) GetByToken(token string) (error, *models.AccountVerification) {
    av := models.AccountVerification{}
    log.Infof("Getting Account verification token: %s", token)
    if store.Where("Token = ?", token).First(&av).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= av.ExpiresAt.Unix() {
        return errors.New(errors.Expired, av.Token), nil
    }

    return nil, &av
}

func (store *accountVerificationStore) New(uid int64) *models.AccountVerification {
    log.Infof("Creating new account verification token for: %d", uid)

    av := &models.AccountVerification{
        UID: uid,
    }

    store.Create(av)

    return av
}
