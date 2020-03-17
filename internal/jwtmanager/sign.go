package jwtmanager

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

func (m *jwtManager) Sign(user *User) (string, error) {
    claims := Claims{
        User: *user,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
            IssuedAt: time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(m.SigningMethod, claims)

    tokenString, err := token.SignedString(m.SignKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
