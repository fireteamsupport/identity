package main

import (
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/profiles/internal/cron"
    "github.com/fireteamsupport/profiles/internal/config"
    "github.com/fireteamsupport/profiles/internal/database"
    "github.com/fireteamsupport/profiles/internal/natsclient"
    "github.com/fireteamsupport/profiles/internal/restserver"
)

var (
    log = logging.New()
    cfg = config.New()
)

func main() {
    log.Info("Starting Account Management for Fireteamsupport...")

    dbClient, err := database.New(/* TODO */)
    if err != nil {
        log.Fatal(err)
    }

    natsClient, err := natsclient.New(/* TODO */, dbClient)
    if err != nil {
        log.Fatal(err)
    }

    restClient, err := restserver.New(/* TODO */, dbClient, natsClient)
    if err != nil {
        log.Fatal(err)
    }

    cronTasks, err := cron.New(dbClient, natsClient)
    if err != nil {
        log.Fatal(err)
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    log.Info("Shuting down...")
    defer dbClient.Close()
    natsClient.Close()

    ctx, cancel := context.WithTimeout(context.Backgound(), 10*time.Second)
    defer cancel()

    if err = restClient.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
}
