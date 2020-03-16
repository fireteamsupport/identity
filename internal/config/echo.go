package config

import (
    "github.com/caarlos0/env/v6"
)

type EchoConfig struct {
    Host string `env:"HOST" envDefault:"0.0.0.0"`
    Port string `env:"PORT" envDefault:"5000"`
}

func EchoLoad() *EchoConfig {
    cfg := new(EchoConfig)
    if err := env.Parse(cfg); err != nil {
        log.Error(err)
        return cfg
    }

    return cfg
}
