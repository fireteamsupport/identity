package database

func (c *client) NewRefreshToken(uid int64, ip string) *RefreshToken {
    log.Infof("Creating new refresh token for: %d", uid)

    rt := &RefreshToken{
        UID: uid,
        IP: ip,
    }

    c.Create(rt)

    return rt
}
