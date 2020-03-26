package models

import (
    "github.com/fireteamsupport/identity/internal/utils"
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
        Verified  bool
    }
)


func (u *User) TableName() string {
    return "profiles"
}

func (u *User) HashPassword() error {
    if u.Password != "" {
        err, pwd := utils.HashPassword(u.Password)
        if err != nil {
            return err
        }

        u.Password = pwd
    }

    return nil
}

func (u *User) NewPassword(password string) {
    u.Password = password
    u.HashPassword()
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
    node, err := snowflake.NewNode(1)
    if err != nil {
        return err
    }

    u.HashPassword()

    snowflake := node.Generate()

    scope.SetColumn("UID", snowflake.Int64())
    scope.SetColumn("CreatedAt", time.Now().UTC())
    scope.SetColumn("UpdatedAt", time.Now().UTC())
    scope.SetColumn("Verified", false)

    return nil
}

func (u *User) ValidPassword(pwd string) bool {
    return utils.ComparePasswordToHash(u.Password, pwd)
}
