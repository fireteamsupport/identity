package jwtmanager

import (
    "github.com/arturoguerra/go-logging"
    "github.com/dgrijalva/jwt-go"
    "crypto/rsa"
    "io/ioutil"
)

var log = logging.New()

type (
    User struct {
        UID      int64  `json:"uid"`
        Username string `json:"username"`
        Email    string `json:"email"`
        Role     int    `json:"role"`
        Banned   bool   `json:"banned"`
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
