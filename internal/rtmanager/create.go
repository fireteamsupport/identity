package rtmanager

func (rt *rtManager) Create(uid int64, ip string) (error, string) {
    refreshToken := rt.Store.New(uid, ip)
    return nil, refreshToken.Token
}
