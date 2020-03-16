package jwtmanager

import (
    "github.com/dgrijalva/jwt-go"
    "crypto/rsa"
    "io/ioutil"
)

type (
    User struct {
        UID      int64
        Username string
        Email    string
        Role     int
    }

    Claims struct {
        User
        jwt.StandardClaims
    }

    jwtManager struct {
        SigningMethod jwt.SigningMethod
        VerifyKey *rsa.PublicKey
        SignKey *rsa.PrivateKey
    }

    JWTManager interface {
        Sign(*User) (string, error)
        Decrypt(string) (error, *Claims)
    }
)

func New(cfg *Config) (error, JWTManager) {
    signBytes, err := ioutil.ReadFile(cfg.PrivKeyPath)
    if err != nil {
        return err, nil
    }

    signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
    if err != nil {
        return err, nil
    }

    verifyBytes, err := ioutil.ReadFile(cfg.PubKeyPath)
    if err != nil {
        return err, nil
    }

    verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
    if err != nil {
        return err, nil
    }

    return nil, &jwtManager{
        SigningMethod: jwt.GetSigningMethod("RS256"),
        VerifyKey: verifyKey,
        SignKey: signKey,
    }
}
