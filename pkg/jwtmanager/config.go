package jwtmanager

import "github.com/caarlos0/env/v6"

type (
    Config struct {
        PrivKeyPath string `env:"JWT_PRIV_KEY_PATH" envDefault:"/etc/keys/rsa"`
        PubKeyPath string `env:"JWT_PUB_KEY_PATH" envDefault:"/etc/keys/rsa.pub"`
    }
)

func NewEnvCfg() (error, *Config) {
    cfg := new(Config)
    if err := env.Parse(cfg); err != nil {
        return err, nil
    }

    return nil, cfg
}
