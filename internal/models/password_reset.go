package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type (
    PasswordReset struct {
        gorm.Model
        Token     string `gorm:"primary_key;auto_increment:false"`
        UID       int64
        ExpiresAt *time.Time
    }
)

func (pr *PasswordReset) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("ExpiresAt", time.Now().Add(time.Hour).UTC().Add(time.Hour))
    token := genToken()
    scope.SetColumn("Token", token)
    return nil
}
