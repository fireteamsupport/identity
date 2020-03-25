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

    return nil, string(hash)
}

func ComparePasswordToHash(hashed, plain string) bool {
    bhash := []byte(hashed)
    bplain := []byte(plain)

    if err := bcrypt.CompareHashAndPassword(bhash, bplain); err != nil {
        return false
    }

    return true
}
