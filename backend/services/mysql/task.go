package mysql

import (
	"database/sql"
	"time"
)

type (
	TaskServiceImpl struct {
		DB *sql.DB
	}
	Task struct {
		ID      uint32
		User_id string `json:"userid"`
		Content string `json:"content"`
		Comment string `json:"comment"`
		Started time.Time
		Ended   time.Time
		Stoped  time.Time
		Status  uint8 `json:"status"`
	}
)

const (
	sqlCreateDatabase = iota
	sqlCreateTable
	sqlInsert
	sqlQuerybyId
	sqlQuerybyUserId
	sqlQuerybyStatus
	sqlModifyStatus
)

var (
	Stop    = 1
	Success = 2

	taskSqls = []string{
		`CREATE DATABASE IF NOT EXISTS task;`,
		`CREATE TABLE IF NOT EXISTS task.task (
			id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			user_id VARCHAR(512) UNSIGNED NOT NULL,
			content VARCHAR(512) NOT NULL DEFAULT ' ',
			comment VARCHAR(512) NOT NULL DEFAULT ' ',
			started DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			ended DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			stoped DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			status TINYINT NOT NULL DEFAULT '0',
			PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INTO task.task (user_id,content,started,ended) VALUES(?,?,?,?)`,
		`SELECT content,comment,started,ended,stoped,status FROM task.task WHERE id = ? LOCK IN SHARE MODE`,
		`SELECT id,content,comment,started,ended,stoped,status FROM task.task WHERE user_id = ? LIMIT (?-1)*10, 10 LOCK IN SHARE MODE`,
		`SELECT id,content,comment,started,ended,stoped FROM task.task WHERE user_id = ? and status = ? LIMIT (?-1)*10, 10 LOCK IN SHARE MODE`,
		`UPDATE task.task SET status = ? , comment = ? , stoped = NOW() WHERE id = ? LIMIT 1`,
	}
)

//Initialize
func (ts *TaskServiceImpl) Initialize() error {
	_, err := ts.DB.Exec(taskSqls[sqlCreateDatabase])
	if err != nil {
		return err
	}

	_, err = ts.DB.Exec(taskSqls[sqlCreateTable])
	if err != nil {
		return err
	}
	return nil
}

//Create task
func (ts *TaskServiceImpl) Create(user_id, content string, started, ended time.Time) error {
	result, err := ts.DB.Exec(taskSqls[sqlInsert], user_id, content, started, ended)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}
	return nil
}

//Query one by id
func (ts *TaskServiceImpl) QueryById(id uint32) (*Task, error) {
	var (
		task Task
	)

	rows, err := ts.DB.Query(taskSqls[sqlQuerybyId], id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&task.Content, &task.Comment, &task.Started, &task.Ended, &task.Stoped, &task.Status); err != nil {
			return nil, err
		}
	}

	return &task, nil
}

//Query by only userid
func (ts *TaskServiceImpl) QueryByUserId(user_id string, page int) ([]*Task, error) {
	var (
		id      uint32
		content string
		comment string
		started time.Time
		ended   time.Time
		stoped  time.Time
		status  uint8

		tasks []*Task
	)

	rows, err := ts.DB.Query(taskSqls[sqlQuerybyUserId], user_id, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&content, &comment, &started, &ended, &stoped, &status); err != nil {
			return nil, err
		}

		task := &Task{
			ID:      id,
			User_id: user_id,
			Content: content,
			Comment: comment,
			Started: started,
			Ended:   ended,
			Stoped:  stoped,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

//query from user's status
func (ts *TaskServiceImpl) QueryByStatus(user_id string, status uint8, page int) ([]*Task, error) {
	var (
		id      uint32
		content string
		comment string
		started time.Time
		ended   time.Time
		stoped  time.Time

		tasks []*Task
	)

	rows, err := ts.DB.Query(taskSqls[sqlQuerybyStatus], user_id, status, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &content, &comment, &started, &ended, &stoped); err != nil {
			return nil, err
		}

		task := &Task{
			ID:      id,
			User_id: user_id,
			Content: content,
			Comment: comment,
			Started: started,
			Ended:   ended,
			Stoped:  stoped,
			Status:  status,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

//stop the task
func (ts *TaskServiceImpl) Stop(comment string, id uint32) error {
	result, err := ts.DB.Exec(taskSqls[sqlModifyStatus], Stop, comment)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}

	return nil
}

//success the task
func (ts *TaskServiceImpl) Success(comment string, id uint32) error {
	result, err := ts.DB.Exec(taskSqls[sqlModifyStatus], Success, comment)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errQueryFailed
	}

	return nil
}
