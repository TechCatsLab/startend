package services

import (
	"database/sql"

	"github.com/TechCatsLab/startend/backend/services/mysql"
)

var (
	// Service -
	TaskService *mysql.TaskServiceImpl
	UserService *mysql.UserServiceImpl
)

// Load all services.
func Load(db *sql.DB) error {
	TaskService = &mysql.TaskServiceImpl{
		DB: db,
	}

	UserService = &mysql.UserServiceImpl{
		DB: db,
	}

	if err := TaskService.Initialize(); err != nil {
		panic(err)
	}

	if err := UserService.Initialize(); err != nil {
		panic(err)
	}

	return nil
}
