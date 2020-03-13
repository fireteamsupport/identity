package cron

import (
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/natsclient"

/*
CRON
- Deletes old refresh tokens
*/

const (
    Weekly = 60 * 60 * 24 * 7
    Daily  = 60 * 60 * 24
    Hourly = 60 * 60
)

func New(dbclient *database.Client) error {
}
