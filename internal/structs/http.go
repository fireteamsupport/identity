package structs

type (
    ReqLogin struct {
        Email    string `json:"email" validate:"required"`
        Password string `json:"password" validate:"required"`
    }

    RespLogin struct {
        AccessToken  string `json:"access_token" validate:"required"`
        RefreshToken string `json:"refresh_token" validate:"required"`
        TokenType    string `json:"token_type" validate:"required"`
    }
)

type (
    ReqRegister struct {
        Username string `json:"username" validate:"required"`
        Email    string `json:"email" validate:"required"`
        Password string `json:"password" validate:"required"`
    }

    RespRegister struct {
        Username string `json:"username" validate:"required"`
        Email    string `json:"email" validate:"required"`
    }
)

type (
    ReqRefresh struct {
        Token string `json:"token" validate:"required"`
    }

    RespRefresh struct {
        AccessToken  string `json:"access_token" validate:"required"`
        RefreshToken string `json:"refresh_token" validate:"required"`
        TokenType    string `json:"token_type" validate:"required"`
    }
)

type (
    ReqLogout struct {
        Token string `json:"token" validate:"required"`
    }
)

type (
    ReqReverify struct {
        Email string `json:"email" validate:"required"`
    }
)

type (
    ReqPasswordReset struct {
        Email string `json:"email" validate:"required"`
    }
)

type (
    ReqRecoverAccount struct {
        Code string `json:"code" validate:"required"`
        Password string `json:"password" validate:"required"`
    }
)

type (
    Empty struct {}
)
