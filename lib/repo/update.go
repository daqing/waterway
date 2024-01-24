package repo

func UpdateFields[T TableNameType](id IdType, fields []KVPair) bool {
	var t T

	tx := DB().Table(t.TableName()).Where("id = ?", id).Updates(buildCondQuery(fields))

	return tx.RowsAffected == 1
}

func UpdateRow[T TableNameType](cond []KVPair, field string, value any) bool {
	var t T

	tx := DB().Table(t.TableName()).Where(buildCondQuery(cond)).Update(field, value)

	return tx.RowsAffected == 1
}
