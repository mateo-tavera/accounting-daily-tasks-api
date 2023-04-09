package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"

	"github.com/mateo-tavera/accounting-daily-tasks/entity"
)

// Global variables
var dateFormatter = "2006-01-02 15:04:05"
var validTechnicians = map[string]string{"tech1": "password1", "tech2": "password2", "tech3": "password2"} //TODO: get from db
var validManagers = map[string]string{"mgr1": "password1", "mgr2": "password2", "mgr3": "password2"}       //TODO: get from db

type RepositoryService interface {
	CreateTask(entity.Task) (int, error)
	GetTask(int, string) (*entity.Task, error)
	GetAllTasks(int, string) ([]entity.Task, error)
	UpdateTask(entity.Task) error
	DeleteTask(int, string) error
}

type repository struct {
	db *sql.DB
}

// Interface implementation
func NewRepository(db *sql.DB) RepositoryService {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetTask(id int, user string) (*entity.Task, error) {

	query := `SELECT id, user, summary, date, status FROM tasks WHERE id = ?`

	//Execute query
	result := new(entity.Task)
	err := repo.db.QueryRow(query, id).Scan(&result.Id, &result.User, &result.Summary, &result.Date, &result.Status)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}

	// Validate authorization
	if user != result.User {
		return nil, fmt.Errorf("permission denied: user %s does not have access to task %d", user, id)
	}

	return result, err
}

func (repo *repository) GetAllTasks(id int, user string) ([]entity.Task, error) {
	// Input parameters
	var tasks []entity.Task
	var rows *sql.Rows
	var err error

	// Validate authorization
	_, isTechnician := validTechnicians[user]
	_, isManager := validManagers[user]
	if !isTechnician && !isManager {
		return nil, errors.New("permission denied")
	}

	//Execute query based on user
	if isManager { // TODO: must verify if user is authorized and registered in repository
		rows, err = repo.db.Query("SELECT id, user, summary, date, status FROM tasks")
	} else {
		rows, err = repo.db.Query("SELECT id, user, summary, date, status FROM tasks WHERE user = ?", user)
	}

	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Parse results from query
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.Id, &task.User, &task.Summary, &task.Date, &task.Status)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}

func (repo *repository) CreateTask(req entity.Task) (int, error) {
	// Validate authorization
	_, isTechnician := validTechnicians[req.User]
	if !isTechnician {
		return 0, errors.New("permission denied")
	}

	// Create validator
	validate := validator.New()

	// Validate according to entity tags
	err := validate.Struct(req)
	if err != nil {
		return 0, fmt.Errorf("error validating input %v", err)
	}

	query := "INSERT INTO tasks (user, summary, date, status) VALUES (?, ?, ?, ?)"

	// Execute query
	result, err := repo.db.Exec(query, req.User, req.Summary, req.Date.Format(dateFormatter), req.Status)
	if err != nil {
		return 0, fmt.Errorf("error executing query: %v", err)
	}

	// Get id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting id: %v", err)
	}
	return int(id), nil
}

func (repo *repository) UpdateTask(task entity.Task) error {
	// Validate authorization
	_, isTechnician := validTechnicians[task.User]
	if !isTechnician {
		return errors.New("permission denied")
	}

	// To be continued
	query := "UPDATE tasks SET date=?," // Date always updates when request is processed

	// Add query params
	var queryParams []interface{}
	queryParams = append(queryParams, time.Now())
	if task.Summary != "" {
		query += " summary=?,"
		queryParams = append(queryParams, task.Summary)
	}
	if task.Status != "" {
		query += " status=?,"
		queryParams = append(queryParams, task.Status)
	}

	// Take off comma and add the id parameter to the query
	query = query[:len(query)-1] + " WHERE id=?"
	queryParams = append(queryParams, task.Id)

	// Execute query
	statement, err := repo.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(queryParams...)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}

func (repo *repository) DeleteTask(id int, user string) error {

	// Validate authorization
	_, isManager := validManagers[user]
	if !isManager {
		return errors.New("permission denied")
	}

	query := "DELETE FROM tasks WHERE id = ?"

	//Prepare statement
	statement, err := repo.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}
	defer statement.Close()
	//Execute query
	_, err = statement.Exec(id)

	return err

}
