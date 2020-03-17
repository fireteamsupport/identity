package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/config"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var log = logging.New()

type (
    client struct {
        *gorm.DB
        Config *config.DBConfig

    }

    Client interface {
        UserLogin(string) (error, *User)
        GetUser(int64) (error, *User)
        RegisterUser(string, string, string) (error, *User)

        NewRefreshToken(int64, string) *RefreshToken
        ClearRefreshTokens(int64) error
        DeleteRefreshToken(string) error
        GetRefreshToken(string) (error, *RefreshToken)
        GetRefreshTokens(int64) (error, []*RefreshToken)

        Save(interface{}) *gorm.DB
        Close() error
    }
)

func connect(username, password, host, dbname string) (*gorm.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)
    db, err := gorm.Open("mysql", connectionString)
    return db, err
}

func (c *client) Init() error {
    c.AutoMigrate(&User{})
    c.AutoMigrate(&RefreshToken{})
    return nil
}

func New(cfg *config.DBConfig) (error, Client) {
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
