package database

import (
    "strconv"
    "time"
    "io"
    "crypto/md5"
)

func genToken() string {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	h := md5.New()
	io.WriteString(h, time)
    return string(h.Sum(nil))
}
