package utils

import (
    "github.com/caarlos0/env/v6"
)

func EnvLoad(i interface{}) (error) {
    if err := env.Parse(i); err != nil {
        return err
    }

    return nil
}
