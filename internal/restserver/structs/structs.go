package structs

// /auth/login
type (
    ReqLogin struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    RespLogin struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
    }
)

// /auth/logout
type (
    ReqLogout struct {
        RefreshToken string
    }
)

// GET /auth/register
type (
    ReqRegister struct {
        Username string
        Email    string
        Password string
    }
)

// GET /auth/refresh
type (
    ReqRefresh struct {
        RefreshToken string
    }

    RespRefresh struct {
        JWT          string
        RefreshToken string
    }
)

// GET /auth/passwordreset
type (
    ReqPasswordReset struct {
        Email string
    }
)

// GET /auth/recover
type (
    ReqRecover struct {
        ResetToken string
        Password   string
    }
)

// GET-PATCH /users/me
type (
    RespGetMe struct {
        Username string
        Email    string
    }

    RespPatchMe struct {
        Username    string
        Email       string
        OldPassword string
        NewPassword string
    }
)

type User struct {
    UID      int64   `json:"uid"`
    Username string  `json:"username"`
    Email    string  `json:"email"`
    Role     int     `json:"role"`
    Banned   bool    `json:"banned"`
}

type (
    Empty struct {}

    Message struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
    }

    HttpError struct {
        Code int `json:"code"`
        Message string `json:"message"`
    }
)
