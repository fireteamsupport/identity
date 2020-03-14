package restutils

import (
    "fmt"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/rtmanager"
    "github.com/fireteamsupport/identity/pkg/jwtmanager"
)

type Options struct {
    Host string
    Port string
    JWTMgmt *jwtmanager.JWTManager
    RTMgmt *rtmanager.RTManager
    DB  database.Client
}

func (opts *Options) GetAddr() string {
    if opts.Host == "" {
        opts.Host = "0.0.0.0"
    }

    if opts.Port == "" {
        opts.Port = "5000"
    }

    fmt.Sprintf("%s:%s", opts.Host, opts.Port)
}
