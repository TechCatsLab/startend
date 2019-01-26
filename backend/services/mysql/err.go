package mysql

import "errors"

var (
	errQueryFailed  = errors.New("affected 0 rows")
	errRowsNotExits = errors.New("there is no such data in database")
)
