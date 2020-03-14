package database


func (c *client) GetRefreshTokens(uid int64) (error, []*RefreshToken) {
    rts := make([]*RefreshToken, 0)
    log.Infof("Getting Refresh tokens for: %d", uid)
    c.Where("UID = ?", uid).First(&rts)
    return nil, rts
}
