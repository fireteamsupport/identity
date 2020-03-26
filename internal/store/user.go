package store

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/fireteamsupport/identity/internal/errors"
    "github.com/fireteamsupport/identity/internal/models"
    "github.com/fireteamsupport/identity/internal/database"
)

type (
    userStore struct {
        *database.Client
    }

    UserStore interface {
        GetId(int64) (error, *models.User)
        GetEmail(string) (error, *models.User)
        New(string, string, string) (error, *models.User)
        Save(interface{}) *gorm.DB
        Delete(interface{}, ...interface{}) *gorm.DB
    }
)

func NewUserStore(db *database.Client) (UserStore, error) {
    db.AutoMigrate(&models.User{})
    return &userStore{db}, nil
}

func (store *userStore) GetId(uid int64) (error, *models.User) {
    u := models.User{}
    log.Infof("Getting user for UID: %d", uid)
        if store.Where("UID = ?", uid).First(&u).RecordNotFound() {
        return errors.New(errors.NotFound, uid), nil
    }

    return nil, &u
}

func (store *userStore) GetEmail(email string) (error, *models.User) {
    u := models.User{}
    log.Infof("Getting user for Email: %s", email)
        if store.Where("Email = ?", email).First(&u).RecordNotFound() {
        return errors.New(errors.NotFound, email), nil
    }

    return nil, &u

}

func (store *userStore) New(username, email, password string) (error, *models.User) {
    u := models.User{}

    if store.Where("Email = ?", email).First(&u).RecordNotFound() {
        newUser := &models.User{
            Username: username,
            Email: email,
            Password: password,
        }

        store.Create(newUser)

        return nil, newUser
    }

    err := errors.New(errors.Exists, email)

    return err, nil
}
