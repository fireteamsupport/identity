package validation

import "github.com/fireteamsupport/identity/internal/utils"

type Config struct {
    APIKey string `json:"email-api-key" env:"EMAIL_API_KEY"`
}

func NewEnvConfig() (err error, cfg *Config) {
    if err := utils.EnvLoad(cfg); err != nil {
        return err, nil
    }

    return
}
