package restutils

import (
    "fmt"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/rtmanager"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/fireteamsupport/identity/internal/email"
)

type Options struct {
    Host string
    Port string
    JWTMgmt jwtmanager.JWTManager
    RTMgmt  rtmanager.RTManager
    DB  database.Client
    Email email.Email
}

func (opts *Options) GetAddr() string {
    if opts.Host == "" {
        opts.Host = "0.0.0.0"
    }

    if opts.Port == "" {
        opts.Port = "5000"
    }

    return fmt.Sprintf("%s:%s", opts.Host, opts.Port)
}
