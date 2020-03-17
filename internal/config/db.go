package config

import (
    "github.com/caarlos0/env/v6"
)

type DBConfig struct {
    Host string `env:"DB_HOST,required"`
    Post string `env:"DB_PORT,required"`
    Name string `env:"DB_NAME,required"`
    User string `env:"DB_USER,required"`
    Password string `env:"DB_PASSWORD,required"`
}

func DBLoad() *DBConfig {
    cfg := new(DBConfig)
    if err := env.Parse(cfg); err != nil {
        log.Error(err)
        return cfg
    }

    return cfg
}
