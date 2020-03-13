package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
    client struct {
        *grom.DB
    }

    Client interface {
        UserLogin(string) (error, *User)
        GetUser(int64) (error, *User)
        RegisterUser(string, string, string) (error, *User)
        NewRefreshToken(int64, string) *RefreshToken
        Save(interface{}) error
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

func (c *client) Close() error {
    return c.DB.Close()
}

func New(username, password, host, dbname string) (Client, error) {
    db, err := connect(username, password, host, dbname)
    if err != nil {
        return nil, err
    }

    c := &client{db}

    if err = c.Init(); err != nil {
        return nil, err
    }

    return c, err
}
