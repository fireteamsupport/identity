package database

func (c *client) NewPasswordReset(uid int64) *PasswordReset {
    log.Infof("Creating new password reset for: %d", uid)

    ps := &PasswordReset{
        UID: uid,
    }

    c.Create(ps)

    return ps
}
