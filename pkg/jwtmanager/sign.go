package jwtmanager

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

func (m *jwtManager) Sign(user *User) (string, error) {
    claims := JWTClaims{
        User: *user,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
            IssuedAt: time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString([]byte(m.Secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
