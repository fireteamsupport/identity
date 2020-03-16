package main

import (
    "github.com/arturoguerra/go-logging"
    "github.com/fireteamsupport/identity/internal/config"
    "github.com/fireteamsupport/identity/internal/database"
    "github.com/fireteamsupport/identity/internal/restserver"
    "github.com/fireteamsupport/identity/internal/restserver/utils"
)

var (
    log = logging.New()
    cfg = config.Load()
)

func main() {
    log.Info("Starting Account Management for Fireteamsupport...")

    err, jwtCfg := jwtmanager.NewEnvCfg()
    if err != nil {
        log.Fatal(err)
    }

    err, jwtManager := jwtmanager.New(jwtCfg)
    if err != nil {
        log.Fatal(err)
    }

    err, rtManager := rtManager.New(cfg.RTCfg)
    if err != nil {
        log.Fatal(err)
    }

    dbClient, err := database.New(cfg.DBCfg)
    if err != nil {
        log.Fatal(err)
    }

    restOpts := restutils.Options{
        DB: dbClient,
        JWTMgmt: jwtManager,
        RTMgmt: rtManager,
    }

    restClient, err := restserver.New(cfg.EchoCfg, opts)
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
