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

func NewDefault() (error, Validate) {
    err, cfg := NewEnvConfig()
    if err != nil {
        return err, nil
    }

    err, v := New(cfg)
    return err, v
}

func New(cfg *Config) (error, Validate) {
    v := &validate{
        Validate: validator.New(),
        Config: cfg,
    }

    log.Info("Registering Email Validation")
    v.RegisterValidation("email", emailValidation(v))

    return nil, v
}
