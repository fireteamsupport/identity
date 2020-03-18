package database

func (c *client) NewAccountVerification(uid int64) *AccountVerification {
    log.Infof("Creating new account verification token for: %d", uid)

    av := &AccountVerification{
        UID: uid,
    }

    c.Create(av)

    return av
}
