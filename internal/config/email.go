package config

import (
    "github.com/caarlos0/env/v6"
)


type EmailConfig struct {
    Region          string `env:"AWS_REGION,required"`
    Sender          string `env:"EMAIL_SENDER,required"`
    AcessKeyID      string `env:"AWS_ACCESS_KEY_ID,required"`
    SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required"`
}

func EmailLoad() (error, *EmailConfig) {
    cfg := new(EmailConfig)
    if err := env.Parse(cfg); err != nil {
        log.Error(err)
        return err, nil
    }

    return nil, cfg
}
