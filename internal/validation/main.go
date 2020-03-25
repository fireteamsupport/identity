package validation

import (
)

var (
    log = logging.New()
)

type (
    Config struct {
        APIKey string `json:"email-api-key" env:"EMAIL_API_KEY"`
    }

    validation struct {
        *validator.Validate
        Config *config.Validation
    }

    Validation struct {
        Struct(interface{}) error

    }
)

func New(cfg *Config) (error, Validation) {
    v := &validation{
        Validate: validator.New(),
        Config: cfg,
    }

    log.Info("Registering Email Validation")
    v.RegisterValidation("email", emailValidation(v))

    return nil, v
}
