package repository_test

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	entity_mocks "github.com/mateo-tavera/accounting-daily-tasks/entity/mocks"
	"github.com/mateo-tavera/accounting-daily-tasks/repository"
	"github.com/stretchr/testify/assert"
)

func DBConncetionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetTask(t *testing.T) {
	// Mock database connection
	database, mock := DBConncetionMock()
	repoMock := repository.NewRepository(database)
	defer database.Close()

	t.Run("error 'Cannot execute query' - GetTask", func(t *testing.T) {
		// Define query and parameters
		dataOk := entity_mocks.GenerateSvcTaskOK()
		query := `SELECT id, user, summary, date, status FROM tasks WHERE id = ?`

		//Expectation mock
		mock.ExpectQuery(query).WithArgs(dataOk.Id).WillReturnError(errors.New("error executing query"))

		// Evaluate results
		response, err := repoMock.GetTask(1, "tech1")
		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("error 'Wrong user' - GetTask", func(t *testing.T) {
		// Define query and parameters
		dataOk := entity_mocks.GenerateSvcTaskOK()
		query := `SELECT id, user, summary, date, status FROM tasks WHERE id = ?`

		//Map dataOk into sql.Row type
		rows := sqlmock.NewRows([]string{"id", "user", "summary", "date", "status"}).
			AddRow(dataOk.Id, "tech2", dataOk.Summary, dataOk.Date, dataOk.Status)

		//Expectation mock
		mock.ExpectQuery(query).WithArgs(dataOk.Id).WillReturnRows(rows)

		// Evaluate results
		response, err := repoMock.GetTask(1, "tech1")
		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("error 'Query executed succesfully' - GetTask", func(t *testing.T) {
		// Define query and parameters
		dataOk := entity_mocks.GenerateSvcTaskOK()
		query := `SELECT id, user, summary, date, status FROM tasks WHERE id = ?`

		// Map dataOk into sql.Row type
		rows := sqlmock.NewRows([]string{"id", "user", "summary", "date", "status"}).
			AddRow(dataOk.Id, dataOk.User, dataOk.Summary, dataOk.Date, dataOk.Status)

		// Expectation mock
		mock.ExpectQuery(query).WithArgs(dataOk.Id).WillReturnRows(rows)

		// Evaluate results
		response, err := repoMock.GetTask(1, "tech1")
		assert.NotNil(t, response)
		assert.Nil(t, err)
	})

}

func TestCreateTask(t *testing.T) {
	// Mock database connection
	database, mock := DBConncetionMock()
	repoMock := repository.NewRepository(database)
	defer database.Close()

	t.Run("error 'Permission denied' - CreateTask", func(t *testing.T) {
		// Define query and parameters
		dataNOk := entity_mocks.GenerateCreateTaskNOK()

		// Evaluate result
		id, err := repoMock.CreateTask(dataNOk)
		assert.Equal(t, id, 0)
		assert.EqualError(t, err, "permission denied")
	})

	t.Run("error 'Cannot execute query' - CreateTask", func(t *testing.T) {
		// Define query and parameters
		query := "INSERT INTO tasks (user, summary, date, status) VALUES (?, ?, ?, ?)"
		dataOk := entity_mocks.GenerateSvcTaskOK()

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(dataOk.Id)

		mock.ExpectQuery(query).WithArgs(dataOk.User, dataOk.Summary, dataOk.Date, dataOk.Status).WillReturnRows(rows)

		id, err := repoMock.CreateTask(entity_mocks.GenerateCreateTaskOK())
		assert.Equal(t, id, 0)
		assert.Error(t, err)
	})

}

func TestDeleteTask(t *testing.T) {
	// Mock database connection
	database, _ := DBConncetionMock()
	repoMock := repository.NewRepository(database)
	defer database.Close()

	t.Run("error 'Permission denied' - DeleteTask", func(t *testing.T) {
		// Define query and parameters
		dataNOk := entity_mocks.GenerateCreateTaskNOK()

		// Evaluate result
		err := repoMock.DeleteTask(dataNOk.Id, dataNOk.User)
		assert.EqualError(t, err, "permission denied")
	})
}
