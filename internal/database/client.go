package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/arturoguerra/go-logging"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/fireteamsupport/identity/internal/models"
)

var log = logging.New()

type (
    client struct {
        *gorm.DB
        Config *Config

    }

    Client interface {
        UserLogin(string) (error, *models.User)
        GetUser(int64) (error, *models.User)
        RegisterUser(string, string, string) (error, *models.User)

        NewRefreshToken(int64, string) *models.RefreshToken
        ClearRefreshTokens(int64) error
        DeleteRefreshToken(string) error
        GetRefreshToken(string) (error, *models.RefreshToken)
        GetRefreshTokens(int64) (error, []*models.RefreshToken)

        NewPasswordReset(int64) *models.PasswordReset
        GetPasswordReset(string) (error, *models.PasswordReset)

        NewAccountVerification(int64) *models.AccountVerification
        GetAccountVerification(string) (error, *models.AccountVerification)

        Save(interface{}) *gorm.DB
        Close() error
        Delete(interface{}, ...interface{}) *gorm.DB
    }
)

func connect(username, password, host, dbname string) (*gorm.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)
    db, err := gorm.Open("mysql", connectionString)
    return db, err
}

func (c *client) Init() error {
    c.AutoMigrate(&models.User{})
    c.AutoMigrate(&models.RefreshToken{})
    c.AutoMigrate(&models.AccountVerification{})
    c.AutoMigrate(&models.PasswordReset{})
    return nil
}

func New(cfg *Config) (error, Client) {
    db, err := connect(cfg.User, cfg.Password, cfg.Host, cfg.Name)
    if err != nil {
        return err, nil
    }

    c := &client{db,cfg}

    if err = c.Init(); err != nil {
        return err, nil
    }

    return err, c
}

func NewDefault() (error, Client) {
    err, cfg := NewEnvConfig()
    if err != nil {
        return err, nil
    }

    err, client := New(cfg)
    return err, client
}
