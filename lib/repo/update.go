package repo

func UpdateFields[T TableNameType](id IdType, fields []KVPair) bool {
	var t T

	db, err := DB()
	if err != nil {
		return false
	}

	tx := db.Table(t.TableName()).Where("id = ?", id).Updates(buildCondQuery(fields))

	return tx.RowsAffected == 1
}

func UpdateRow[T TableNameType](cond []KVPair, field string, value any) bool {
	var t T

	db, err := DB()
	if err != nil {
		return false
	}

	tx := db.Table(t.TableName()).Where(buildCondQuery(cond)).Update(field, value)

	return tx.RowsAffected == 1
}
