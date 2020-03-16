package rtmanager

func (rt *rtManager) GetAll(uid int64) (error, []*RToken) {
    err, rawrts := rt.DB.GetRefreshTokens(uid)
    if err != nil {
        return err, nil
    }

    rts := make([]*RToken, 0)
    for _, item := range rawrts {
        rts = append(rts, &RToken{
            UID: item.UID,
            Token: item.Token,
            ExpiresAt: item.ExpiresAt.Unix(),
            IP: item.IP,
        })
    }

    return nil, rts
}
