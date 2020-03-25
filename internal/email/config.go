package email

import "github.com/fireteamsupport/identity/internal/utils"


type Config struct {
    Region          string `env:"AWS_REGION,required"`
    Sender          string `env:"EMAIL_SENDER,required"`
    AcessKeyID      string `env:"AWS_ACCESS_KEY_ID,required"`
    SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required"`
}

func NewEnvConfig() (error, *Config) {
    c := new(Config)
    if err := utils.EnvLoad(c); err != nil {
        return err, nil
    }

    return nil, c
}
