package utils

import (
    "strings"
    "reflect"
    "github.com/go-playground/validator/v10"
)

func Validator() *validator.Validate {
    v := validator.New()

    v.RegisterTagNameFunc(func(fld reflect.StructField) string {
        n := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if n == "-" {
            return ""
        }

        return n
    })

    return v
}
