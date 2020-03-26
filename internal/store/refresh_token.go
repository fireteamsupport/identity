package store

import (
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
    "github.com/fireteamsupport/identity/internal/database"
)

type (
    refreshTokenStore struct {
        *database.Client
    }

    RefreshTokenStore interface {
        New(int64, string) *models.RefreshToken
        GetByToken(string) (error, *models.RefreshToken)
        GetAllByUID(int64) (error, []*models.RefreshToken)
        DeleteByToken(string) error
        ClearByUID(int64) error

        Save(interface{}) *gorm.DB
        Delete(interface{}, ...interface{}) *gorm.DB
    }
)

func NewRefreshTokenStore(db *database.Client) (RefreshTokenStore, error) {
    db.AutoMigrate(&models.RefreshToken{})
    return &refreshTokenStore{db}, nil
}

func (store *refreshTokenStore) New(uid int64, ip string) *models.RefreshToken {
    log.Infof("Creating new refresh token for: %d", uid)

    rt := &models.RefreshToken{
        UID: uid,
        IP: ip,
    }

    store.Create(rt)

    return rt
}

func (store *refreshTokenStore) GetByToken(token string) (error, *models.RefreshToken) {
    rt := models.RefreshToken{}

    log.Infof("Getting refresh token: %s", token)
    if store.Where("Token = ?", token).First(&rt).RecordNotFound() {
        return errors.New(errors.NotFound, token), nil
    }

    if time.Now().UTC().Unix() >= rt.ExpiresAt.Unix() {
        return errors.New(errors.Expired, rt.Token), nil
    }

    return nil, &rt
}

func (store *refreshTokenStore) GetAllByUID(uid int64) (error, []*models.RefreshToken) {
    rts := make([]*models.RefreshToken, 0)
    log.Info("Getting refresh tokens for: %d", uid)
    store.Where("UID = ?", uid).First(&rts)
    return nil, rts
}

func (store *refreshTokenStore) DeleteByToken(token string) error {
    log.Infof("Deleting refresh token: %s", token)
    store.Where("Token = ?", token).Delete(&models.RefreshToken{})
    return nil
}

func (store *refreshTokenStore) ClearByUID(uid int64) error {
    log.Infof("Clearing all refresh tokens for: %d", uid)
    store.Where("UID = ?", uid).Delete(&models.RefreshToken{})
    return nil
}
