package jwtmanager

import "github.com/fireteamsupport/identity/internal/utils"

type (
    Config struct {
        PrivKeyPath string `env:"JWT_PRIV_KEY_PATH" envDefault:"/etc/keys/rsa"`
        PubKeyPath string `env:"JWT_PUB_KEY_PATH" envDefault:"/etc/keys/rsa.pub"`
    }
)

func NewEnvConfig() (*Config, error) {
    cfg := new(Config)
    if err := utils.EnvLoad(cfg); err != nil {
        return nil, err
    }

    return cfg, nil
}
