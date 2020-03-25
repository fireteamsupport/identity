package main

import (
    "os"
    "time"
    "os/signal"
    "syscall"
    "context"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/email"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/rtmanager"
    "github.com/fireteamsupport/identity/internal/jwtmanager"
    "github.com/fireteamsupport/identity/internal/restserver"
    "github.com/fireteamsupport/identity/internal/validation"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
)

var (
    log = logging.New()
)

func main() {
    log.Info("Starting Account Management for Fireteamsupport...")

    err, dbClient := database.NewDefault()
    log.Info("Starting database..")
    if err != nil {
        log.Fatal(err)
    }

    err, emailClient := email.NewDefault()
    if err != nil {
        log.Fatal(err)
    }

    err, validate := validation.NewDefault()
    log.Info("Starting Struct Validator...")
    if err != nil {
        log.Fatal(err)
    }

    err, jwtManager := jwtmanager.NewDefault()
    log.Info("Starting JWTMananger..")
    if err != nil {
        log.Fatal(err)
    }

    err, rtManager := rtmanager.New(dbClient)
    log.Info("Starting Refresh Token Manager...")
    if err != nil {
        log.Fatal(err)
    }

    restOpts := &restutils.Options{
        DB: dbClient,
        JWTMgmt: jwtManager,
        RTMgmt: rtManager,
        Email: emailClient,
        Validate: validate,
    }

    err, restClient := restserver.NewDefault(restOpts)
    if err != nil {
        log.Fatal(err)
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    log.Info("Shuting down...")
    defer dbClient.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err = restClient.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
}
