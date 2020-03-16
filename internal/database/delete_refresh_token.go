package database


func (c *client) DeleteRefreshToken(token string, uid int64) (error) {
    log.Infof("Deleting Refresh token: %s", token)
    c.Where("Token = ? AND UID = ?", token, uid).Delete(&RefreshToken{})
    return nil
}
