package errors

import (
    "fmt"
)

const (
    BadRequest = 400
    Forbidden = 401
    Unauthorized = 403
    Expired = 403
    NotFound = 404
    Exists = 400
    InternalServerError = 502
)

type (
    JsonError struct {
        Code int `json:"code"`
        Message string `json:"message"`
    }

    ErrorType interface {
        Code() int
        Error() string
        Json() *JsonError
    }

    errorType struct {
        ErrCode int
        Message string
    }
)

func (err *errorType) Code() int {
    return err.ErrCode
}

func (err *errorType) Error() string {
    return fmt.Sprintf("Code: %d Error: %s", err.ErrCode, err.Message)
}

func (err *errorType) Json() *JsonError {
    return &JsonError{
        Code: err.ErrCode,
        Message: err.Message,
    }
}

func New(code int, ctx interface{}) ErrorType {
    return &errorType{
        ErrCode: code,
        Message: fmt.Sprintf("%v", ctx),
    }
}
