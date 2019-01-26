package mysql

import (
	"database/sql"
)

type (
	UserServiceImpl struct {
		DB *sql.DB
	}
)

const (
	sqlCreateUserTable = iota
	sqlInsertUser
	sqlQueryUser
)

var (
	userSqls = []string{
		`CREATE TABLE IF NOT EXISTS task.user (
			uid BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			user_id VARCHAR(512) UNIQUE  NOT NULL,
			PRIMARY KEY (uid)
		) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INTO task.user (user_id) VALUES (?)`,
		`SELECT uid FROM task.user WHERE user_id = ? LOCK IN SHARE MODE`,
	}
)

//Initialize user talbe
func (us *UserServiceImpl) Initialize() error {
	_, err := us.DB.Exec(userSqls[sqlCreateUserTable])
	if err != nil {

		return err
	}

	return nil
}

//create to register user
func (us *UserServiceImpl) create(user_id string) (uint32, error) {
	result, err := us.DB.Exec(userSqls[sqlInsertUser], user_id)
	if err != nil {

		return 0, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {

		return 0, errQueryFailed
	}

	id, err := result.LastInsertId()
	if err != nil {

		return 0, err
	}

	return uint32(id), nil
}

//Query user if register
func (us *UserServiceImpl) QueryUid(user_id string) (uid uint32, err error) {
	err = us.DB.QueryRow(userSqls[sqlQueryUser], user_id).Scan(&uid)
	if err != nil {
		if err.Error() == errNoRows.Error() {
			if uid, err = us.create(user_id); err != nil {
				return
			}
		}
	}

	return
}
