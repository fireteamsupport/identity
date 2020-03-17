package database

import (
    "github.com/bwmarrin/snowflake"
    "github.com/jinzhu/gorm"
    "time"
)

type (
    User struct {
        gorm.Model
        UID       int64  `gorm:"primary_key;auto_increment:false"`
        Email     string `gorm:"type:varchar(100);unique_index"`
        Username  string
        Password  string
        Role      int
        Banned    bool
    }

    RefreshToken struct {
        gorm.Model
        Token     string `gorm:"primary_key;auto_increment:false"`
        UID       int64
        IP        string
        ExpiresAt *time.Time
    }
)

//func (u *User) TableName() string {
//    return "profiles"
//}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
    node, err := snowflake.NewNode(1)
    if err != nil {
        return err
    }

    snowflake := node.Generate()

    scope.SetColumn("UID", snowflake.Int64())
    scope.SetColumn("CreatedAt", time.Now().UTC())
    scope.SetColumn("UpdatedAt", time.Now().UTC())

    return nil
}


func (rf *RefreshToken) TableName() string {
    return "refresh_tokens"
}

func (rf *RefreshToken) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("ExpiresAt", time.Now().Add(336 * time.Hour).Unix())

    token := genToken()

    if err := scope.SetColumn("Token", token); err != nil {
        return err
    }

    scope.SetColumn("CreatedAt", time.Now().UTC())
    scope.SetColumn("ExpiresAt", time.Now().UTC().Add(336 * time.Hour))

    return nil
}
