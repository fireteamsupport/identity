package database


func (c *client) ClearRefreshTokens(uid int64) (error) {
    log.Infof("Clearing all Refresh tokens for: %d", uid)
    c.Where("UID = ?", uid).Delete(&RefreshToken{})
    return nil
}
