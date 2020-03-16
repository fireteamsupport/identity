package structs

type User struct {
    UID      int64 `json:"uid"`
    Email    string `json:"email"`
    Username string `json:"username"`
}
