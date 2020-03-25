package userroutes

type resp_User struct {
    UID      int64  `json:"uid"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Role     int    `json:"role"`
    Banned   bool   `json:"banned"`
    Verified bool   `json:"verified"`
}
