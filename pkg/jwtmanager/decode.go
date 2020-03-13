package jwtmanager

import (
    "github.com/dgrijalva/jwt-go"
    "errors"
    "fmt"
)

func (m *jwtManager) Decrypt(tokenString string) (error, *JWTClaims) {
    claims := new(JWTClaims)

    token, err := jwt.ParseWithClaims(tokenString, claims, func(time *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(m.Secret), nil
    })

    if !token.Valid {
        err = errors.New("Invalid Token")
        return err, nil
    }

    return nil, &claims
}
