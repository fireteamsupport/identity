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

    log.Info("Starting database..")
    err, dbClient := database.NewDefault()
    if err != nil {
        log.Fatal(err)
    }

    log.Info("Starting Temp Email client, Will be moved to its own package and use nats later..")
    err, emailClient := email.NewDefault()
    if err != nil {
        log.Fatal(err)
    }

    log.Info("Starting Struct Validator...")
    err, validate := validation.NewDefault()
    if err != nil {
        log.Fatal(err)
    }

    log.Info("Starting JWTMananger..")
    err, jwtManager := jwtmanager.NewDefault()
    if err != nil {
        log.Fatal(err)
    }

    log.Info("Starting Refresh Token Manager...")
    err, rtManager := rtmanager.New(dbClient)
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
