package main

import (
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/config"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/restserver"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
    "github.com/fireteamsupport/identity/internal/utils"
)

var (
    log = logging.New()
    cfg = config.New()
)

func main() {
    log.Info("Starting Account Management for Fireteamsupport...")

    err, jwtManager := jwtmanager.New(/* JWT Secret */)
    if err != nil {
        log.Fatal(err)
    }

    err, rtManager := rtManager.New(/* Refresh token mananger */)
    if err != nil {
        log.Fatal(err)
    }

    dbClient, err := database.New(/* TODO */)
    if err != nil {
        log.Fatal(err)
    }

    restOpts := restutils.Options{
        Host: cfg.HTTP.Host,
        Port: cfg.HTTP.Port,
        DB: dbClient,
        JWTMgmt: jwtManager,
        RTMgmt: rtManager,
    }

    restClient, err := restserver.New(opts)
    if err != nil {
        log.Fatal(err)
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    log.Info("Shuting down...")
    defer dbClient.Close()

    ctx, cancel := context.WithTimeout(context.Backgound(), 10*time.Second)
    defer cancel()

    if err = restClient.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
}
