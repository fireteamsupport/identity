package structs

type (
    ReqLogin struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    RespLogin struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
        TokenType    string `json:"token_type"`
    }
)


type (
    Empty struct {}
)
