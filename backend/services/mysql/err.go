package mysql

import "errors"

var (
	errQueryFailed = errors.New("affected 0 rows")
	errNoRows      = errors.New("sql: no rows in result set")
)
