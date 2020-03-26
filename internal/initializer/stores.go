package initializer

import (
    "github.com/fireteamsupport/identity/internal/store"
    "github.com/fireteamsupport/identity/internal/database"
)

type Stores struct {
    DB *database.Client
    User store.UserStore
    RefreshToken store.RefreshTokenStore
    PasswordReset store.PasswordResetStore
    AccountVerification store.AccountVerificationStore
}
