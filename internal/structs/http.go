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
    Empty struct {}
)
