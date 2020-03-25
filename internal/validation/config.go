package validation

import "github.com/fireteamsupport/identity/internal/utils"

type Config struct {
    EmailAPIKey string `json:"email-api-key" env:"EMAIL_API_KEY"`
}

func NewEnvConfig() (error, *Config) {
    cfg := new(Config)
    if err := utils.EnvLoad(cfg); err != nil {
        return err, nil
    }

    return nil, cfg
}
