package database

import (
    "github.com/bwmarrin/snowflake"
    "github.com/jinzhu/gorm"
    "time"
    "fmt"
)

type User struct {
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
