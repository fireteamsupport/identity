package jwtmanager

import "github.com/fireteamsupport/identity/internal/utils"

type (
    Config struct {
        PrivKeyPath string `env:"JWT_PRIV_KEY_PATH" envDefault:"/etc/keys/rsa"`
        PubKeyPath string `env:"JWT_PUB_KEY_PATH" envDefault:"/etc/keys/rsa.pub"`
    }
)

func NewEnvConfig() (error, *Config) {
    cfg := new(Config)
    if err := utils.EnvLoad(cfg); err != nil {
        return err, nil
    }

    return nil, cfg
}
