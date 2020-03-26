// +build wireinject

package main

import (
    "github.com/fireteamsupport/identity/internal/initializer"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/store"
    "github.com/fireteamsupport/identity/internal/validation"
    "github.com/fireteamsupport/identity/internal/rtmanager"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/google/wire"
)

var storeSet = wire.NewSet(
    database.NewDefault,
    store.NewUserStore,
    store.NewAccountVerificationStore,
    store.NewPasswordResetStore,
    store.NewRefreshTokenStore,
    wire.Struct(new(initializer.Store), "*"),
)

var restStore = wire.NewSet(
    storeSet,
    jwtmanager.NewDefault,
    rtmanager.New
    validation.NewDefault,
    wire.Struct(new(initializer.Rest), "*"),

func Stores() (*initalizer.Store, error) {
    wire.Build(storeSet)
    return &initializer.Rest{}, nil
}

func Rest() (*initializer.Rest, error) {
    wire.Build(resetSet)
    return &initializer.Rest{}, nil
}
