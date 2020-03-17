package database


func (c *client) DeleteRefreshToken(token string) error {
    log.Infof("Deleting Refresh token: %s", token)
    c.Where("Token = ?", token).Delete(&RefreshToken{})
    return nil
}
