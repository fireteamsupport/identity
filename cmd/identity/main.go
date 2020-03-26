package main

import (
    "os"
    "time"
    "os/signal"
    "syscall"
    "context"
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/restserver"
)

var (
    log = logging.New()
)

func main() {
    log.Info("Starting Account Management for Fireteamsupport...")

    restOpts, err := Rest()
    if err != nil {
        log.Fatal(err)
    }

    log.Info("Starting ECHO Server...")
    restClient, err := restserver.NewDefault(restOpts)
    if err != nil {
        log.Fatal(err)
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    log.Info("Shuting down...")
    defer restOpts.Store.DB.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err = restClient.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
}
