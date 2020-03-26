package initializer

import (
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/fireteamsupport/identity/internal/rtmanager"
    "github.com/fireteamsupport/identity/internal/validation"
    "github.com/fireteamsupport/identity/internal/email"
)

type Rest struct {
    Store *Stores
    JWT   jwtmanager.JWTManager
    Email email.Email
    Validate validation.Validate
    RT    rtmanager.RTManager
}
