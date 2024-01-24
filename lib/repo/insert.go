package repo

func Insert[T TableNameType](attributes []KVPair) (*T, error) {
	return InsertSkipExists[T](attributes, false)
}

func InsertSkipExists[T TableNameType](attributes []KVPair, skipExists bool) (*T, error) {
	if skipExists {
		ex, err := Exists[T](attributes)
		if err != nil {
			return nil, err
		}

		if ex {
			return nil, nil
		}
	}

	var t T

	DB().Table(t.TableName()).Create(buildCondQuery(attributes)).Scan(&t)

	return &t, nil
}
