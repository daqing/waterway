package repo

func Delete[T TableNameType](conds []KVPair) error {
	var t T

	DB().Table(t.TableName()).Where(buildCondQuery(conds)).Delete(&t)

	return nil
}
