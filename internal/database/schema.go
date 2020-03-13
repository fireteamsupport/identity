package database

import (
    "github.com/bwmarrin/snowflake"
    "github.com/jinzhu/gorm"
    "time"
    "fmt"
)

type (
    User struct {
        gorm.Model
        UID          int64  `gorm:"primary_key;auto_increment:false"`
        Email        string `gorm:"type:varchar(100);unique_index"`
        Username     string
        Password     string
        Google       string
        Discord      string
        Admin        bool
        Staff        bool
        Employee     bool
        Banned       bool
        CreationAt   *time.Time
        UpdatedAt   *time.Time
    }

    RefreshToken struct {
        gorm.Model
        Token     string `gorm:"primary_key;auto_increment:false"`
        UID       int64
        ExpiresAt *time.Time
        IP        string
    }
)

func (u *User) TableName() string {
    return "profiles"
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
    node, err := snowflake.NewNode(1)
    if err != nil {
        return err
    }

    snowflake := node.Generate()

    return scope.SetColumn("UID", snowflake.Int64())
}


func (rf *RefreshToken) TableName() string {
    return "refresh_tokens"
}

func (rf *RefreshTOken) BeforeCreate(scope *gorm.Scope) error {
    /* We use UTC as our default time */
    loc, _ := time.LoadLocation("UTC")

    /* Refresh tokens will naturally expire in 2 weeks */
    exptime := time.Now().In(loc).Add(336 * time.Hour)

    if err := scope.SetColumn("ExpiresAt", exptime); err != nil {
        return err
    }

    token := genToken()

    if err = scope.SetColumn("Token", token); err != nil {
        return err
    }

    return nil
}
