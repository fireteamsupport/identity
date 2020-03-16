package rtmanager

func (rt *rtManager) Get(token string) (*RToken, error) {
    err, rtoken := rt.DB.GetRefreshToken(token)
    if err != nil {
        return nil, err
    }

    return &RToken{
        UID: rtoken.UID,
        Token: rtoken.Token,
        IP: rtoken.IP,
    }, nil
}
