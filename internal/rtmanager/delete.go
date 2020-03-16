package rtmanager

func (rt *rtManager) Delete(token string, uid int64) error {
    rt.DB.DeleteRefreshToken(token, uid)
    return nil
}
