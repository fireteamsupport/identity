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
    Client struct {
        *gorm.DB
        Config *Config
    }

)

func connect(username, password, host, dbname string) (*gorm.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)
    db, err := gorm.Open("mysql", connectionString)
    return db, err
}

func (c *Client) Init() error {
    c.AutoMigrate(&models.User{})
    c.AutoMigrate(&models.RefreshToken{})
    c.AutoMigrate(&models.AccountVerification{})
    c.AutoMigrate(&models.PasswordReset{})
    return nil
}

func New(cfg *Config) (*Client, error) {
    db, err := connect(cfg.User, cfg.Password, cfg.Host, cfg.Name)
    if err != nil {
        return nil, err
    }

    c := &Client{db,cfg}

    if err = c.Init(); err != nil {
        return nil, err
    }

    return c, err
}

func NewDefault() (*Client, error) {
    err, cfg := NewEnvConfig()
    if err != nil {
        return nil, err
    }

    client, err := New(cfg)
    return client, err
}
