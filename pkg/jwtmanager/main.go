package jwtmanager

type (
    User struct {
        UID      int64
        Username string
        Email    string
        Role     int
    }

    JWTClaims struct {
        User
        jwt.StandardClaims
    }

    jwtManager struct {
        Secret string
    }

    JWTManager interface {
        Sign(*User) (error, *JWTClaims)
        Decrypt() (string, error)
    }
)


func New(secret string) (error, JWTManager) {
    if secret == "" {
        return errors.New("Invalid secret"), nil
    }

    return nil, &jwtManager{
        Secret: secret,
    }
}
