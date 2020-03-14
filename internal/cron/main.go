package cron

import (
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/events"
)
/*
CRON
- Deletes old refresh tokens
*/

var log = logging.New()

const (
    Weekly = 60 * 60 * 24 * 7
    Daily  = 60 * 60 * 24
    Hourly = 60 * 60
)

func New(dbclient *database.Client, events events.Channels) error {
}
