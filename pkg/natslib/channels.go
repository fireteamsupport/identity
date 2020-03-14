package natslib

import (
    "time"
)

type (
    InvalidateRefreshToken struct {
        UID int64
        Time *time.Time
    }
)
