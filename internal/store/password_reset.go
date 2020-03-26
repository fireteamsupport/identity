package store

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
    "github.com/fireteamsupport/identity/internal/database"
)

type (
    passwordResetStore struct {
        *database.Client
    }

    PasswordResetStore interface {
        New(int64) error
        GetByToken(string) (error, *models.PasswordReset)
        Save(interface{}) *gorm.DB
        Delete(interface{}, ...interface{}) *gorm.DB
    }
)

func NewPasswordResetStore(db *database.Client) (*PasswordResetStore, error) {
    db.AutoMigrate(&models.PasswordReset{})
    return &passwordResetStore{db}, nil
}

func (store *passwordResetStore) GetByToken(token string) (error, *models.PasswordReset) {
    pr := models.PasswordReset{}
    log.Infof("Getting Account verification token: %s", token)
    if store.Where("Token = ?", token).First(&pr).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= pr.ExpiresAt.Unix() {
        return errors.New(errors.Expired, pr.Token), nil
    }

    return nil, &pr
}

func (store *passwordResetStore) New(uid int64) *models.PasswordReset {
    log.Infof("Creating new password reset token for: %d", uid)

    pr := &models.PasswordReset{
        UID: uid,
    }

    store.Create(pr)

    return pr
}
