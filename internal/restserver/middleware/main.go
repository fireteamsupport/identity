package middleware

import (
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/arturoguerra/go-logging"
)

var log = logging.New()

type (
    Middleware struct {
        JWTMgmt jwtmanager.JWTManager
    }
)

func New(jwt jwtmanager.JWTManager) *Middleware {
    return &Middleware{jwt}
}
