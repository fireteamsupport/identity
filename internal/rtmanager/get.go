package rtmanager

func (rt *rtManager) Get(token string) (error, *RToken) {
    err, rtoken := rt.Store.GetByToken(token)
    if err != nil {
        return err, nil
    }

    return nil, &RToken{
        UID: rtoken.UID,
        Token: rtoken.Token,
        IP: rtoken.IP,
    }
}
