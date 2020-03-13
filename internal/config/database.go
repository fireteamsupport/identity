package config

import (
    "github.com/fireteamsupport/identity/internal/structs"
    "os"
)

func LoadDBConfig() *structs.DBConfig {
    return &structs.DBConfig{}
}
