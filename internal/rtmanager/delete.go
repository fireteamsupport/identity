package rtmanager

func (rt *rtManager) Delete(token string) error {
    rt.Store.DeleteByToken(token)
    return nil
}
