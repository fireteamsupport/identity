package database

import (
    "fmt"
    "strconv"
    "time"
    "io"
    "crypto/md5"
)

func getToken() string {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	h := md5.New()
	io.WriteString(h, time)
    return h.Sum(nil)
}
