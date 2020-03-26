package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (error, string) {
    bpwd := []byte(pwd)

    hash, err := bcrypt.GenerateFromPassword(bpwd, bcrypt.MinCost)
    if err != nil {
        return err, ""
    }

    log.Info(string(hash))
    return nil, string(hash)
}

func ComparePasswordToHash(hashed, plain string) bool {
    bhash := []byte(hashed)
    bplain := []byte(plain)

    log.Info(hashed)
    log.Info(plain)

    if err := bcrypt.CompareHashAndPassword(bhash, bplain); err != nil {
        log.Error(err)
        return false
    }

    return true
}
