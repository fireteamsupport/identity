package jwtmanager

import (
    "github.com/dgrijalva/jwt-go"
    "errors"
    "strings"
    "fmt"
)

func (m *jwtManager) Decrypt(tokenString string) (error, *Claims) {
    claims := new(Claims)
    tokenString = strings.TrimSpace(tokenString)

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        log.Infof("%v", token.Method)
        if token.Method != jwt.SigningMethodRS256 {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }


        return m.VerifyKey, nil
    })

    if err != nil {
        return err, nil
    }


    if !token.Valid {
        err = errors.New("Invalid Token")
        return err, nil
    }

    return nil, claims
}
