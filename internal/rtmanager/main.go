package rtmanager


/*
This package handles RefreshToken management
it stores data in the db with the following format:
uid userid int64
token refresh token string
expiresAt timedata
deletedAt same as expires at
*/

import (
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/store"
)

var log = logging.New()

type (
    RToken struct {
        UID       int64
        Token     string
        ExpiresAt int64
        IP        string
    }

    rtManager struct {
        Store store.RefreshTokenStore
    }

    RTManager interface {
        Create(int64, string) (error, string)
        GetAll(int64) (error, []*RToken)
        Get(string) (error, *RToken)
        Delete(string) error
    }
)

func New(store store.RefreshTokenStore) (error, RTManager) {
    return nil, &rtManager{store}
}
