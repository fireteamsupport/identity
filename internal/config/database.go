package config

import (
    "github.com/fireteamsupport/profiles/internal/structs"
    "os"
)

func LoadDBConfig() *structs.DBConfig {
    return &structs.DBConfig{}
}
