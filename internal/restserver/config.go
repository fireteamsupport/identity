package restserver

import "github.com/fireteamsupport/identity/internal/utils"

type Config struct {
    Host string `env:"HOST" envDefault:"0.0.0.0"`
    Port string `env:"PORT" envDefault:"5000"`
}

func NewEnvConfig() (error, *Config) {
    c := new(Config)
    if err := utils.EnvLoad(c); err != nil {
        return err, nil
    }

    return nil, c
}
