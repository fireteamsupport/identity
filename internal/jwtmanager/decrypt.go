package jwtmanager

import (
    "github.com/dgrijalva/jwt-go"
    "errors"
    "fmt"
)

func (m *jwtManager) Decrypt(tokenString string) (error, *Claims) {
    claims := new(Claims)

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return m.VerifyKey, nil
    })

    if !token.Valid {
        err = errors.New("Invalid Token")
        return err, nil
    }

    return nil, claims
}
