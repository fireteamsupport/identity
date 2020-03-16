package rtmanager

func (rt *rtManager) Create(uid int64, ip string) (string, error) {
    refreshToken := rt.DB.NewRefreshToken(uid, ip)
    return refreshToken.Token, nil
}
