package database

import "github.com/fireteamsupport/identity/internal/utils"

type Config struct {
    Host     string `env:"DB_HOST,required"`
    Post     string `env:"DB_PORT,required"`
    Name     string `env:"DB_NAME,required"`
    User     string `env:"DB_USER,required"`
    Password string `env:"DB_PASSWORD,required"`
}

func NewEnvConfig() (error, *Config) {
    c := new(Config)
    if err := utils.EnvLoad(c); err != nil {
        return err, nil
    }

    return nil, c
}
