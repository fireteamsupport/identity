package validation

import (
    "github.com/go-playground/validator/v10"
    "github.com/arturoguerra/go-logging"
)

var log = logging.New()

type (
    validate struct {
        *validator.Validate
        Config *Config
    }

    Validate interface {
        Struct(interface{}) error
    }
)

func NewDefault() (Validate, error) {
    cfg, err := NewEnvConfig()
    if err != nil {
        return nil, err
    }

    v, err := New(cfg)
    return v, err
}

func New(cfg *Config) (Validate, error) {
    v := &validate{
        Validate: validator.New(),
        Config: cfg,
    }

    log.Info("Registering Email Validation")
    v.RegisterValidation("email", emailValidation(v))

    return v, nil
}
