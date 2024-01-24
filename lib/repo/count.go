package repo

func Count[T TableNameType](conds []KVPair) (n int64, err error) {
	var t T

	DB().Table(t.TableName()).Where(buildCondQuery(conds)).Count(&n)

	return n, nil
}
