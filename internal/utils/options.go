package utils

import (
    "github.com/fireteamsupport/identity/pkg/jwtmanager"
    "github.com/fireteamsupport/identity/internal/database"
)

type Options struct {
    DBClient database.Client

    JWT jwtmanager.JWTManager
}

func NewOpts(db database.Client, jwt jwtmanager.JWTManager) *Options {
    return &Options{
        DBClient: db,
        JWT: jwt,
    }
}
