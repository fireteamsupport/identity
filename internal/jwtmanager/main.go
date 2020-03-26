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

func NewDefault() (JWTManager, error) {
    cfg, err := NewEnvConfig()
    if err != nil {
        return nil, err
    }

    jwt, err := New(cfg)
    return jwt, err
}

func New(cfg *Config) (JWTManager, error) {
    signBytes, err := ioutil.ReadFile(cfg.PrivKeyPath)
    if err != nil {
        return nil, err
    }

    signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
    if err != nil {
        return nil, err
    }

    verifyBytes, err := ioutil.ReadFile(cfg.PubKeyPath)
    if err != nil {
        return nil, err
    }

    verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
    if err != nil {
        return nil, err
    }

    return &jwtManager{
        SigningMethod: jwt.GetSigningMethod("RS256"),
        VerifyKey: verifyKey,
        SignKey: signKey,
    }, nil
}
