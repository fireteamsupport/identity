package rtmanager

func (rt *rtManager) Delete(token string) error {
    rt.DB.DeleteRefreshToken(token)
    return nil
}
