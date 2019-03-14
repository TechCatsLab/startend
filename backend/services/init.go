package services

import (
	"database/sql"

	"github.com/TechCatsLab/startend/backend/services/mysql"
)

var (
	// Service -
	TaskService *mysql.TaskServiceImpl
)

// Load all services.
func Load(db *sql.DB) error {
	TaskService = &mysql.TaskServiceImpl{
		DB: db,
	}

	if err := TaskService.Initialize(); err != nil {
		panic(err)
	}

	return nil
}
