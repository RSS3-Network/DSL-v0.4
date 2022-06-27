package database

import "database/sql"

func WrapNullBool(b bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: true,
	}
}

func UnwrapNullBool(b sql.NullBool) bool {
	return b.Valid && b.Bool
}
