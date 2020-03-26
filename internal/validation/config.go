package validation

import "github.com/fireteamsupport/identity/internal/utils"

type Config struct {
    EmailAPIKey string `json:"email-api-key" env:"EMAIL_API_KEY,required"`
}

func NewEnvConfig() (*Config, error) {
    cfg := new(Config)
    if err := utils.EnvLoad(cfg); err != nil {
        return nil, err
    }

    return cfg, nil
}
