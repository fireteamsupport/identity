package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type (
    RefreshToken struct {
        gorm.Model
        Token     string `gorm:"primary_key;auto_increment:false"`
        UID       int64
        IP        string
        ExpiresAt *time.Time
    }
)

func (rf *RefreshToken) TableName() string {
    return "refresh_tokens"
}

func (rf *RefreshToken) BeforeCreate(scope *gorm.Scope) error {

    token := genToken()

    if err := scope.SetColumn("Token", token); err != nil {
        return err
    }

    scope.SetColumn("CreatedAt", time.Now().UTC())
    scope.SetColumn("ExpiresAt", time.Now().UTC().Add(336 * time.Hour))

    return nil
}
