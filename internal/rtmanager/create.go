package rtmanager

func (rt *rtManager) Create(uid int64, ip string) (error, string) {
    refreshToken := rt.DB.NewRefreshToken(uid, ip)
    return nil, refreshToken.Token
}
