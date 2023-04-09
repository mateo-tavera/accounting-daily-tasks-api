package service_test

import (
	"errors"
	"testing"

	"github.com/mateo-tavera/accounting-daily-tasks/entity"
	entity_mocks "github.com/mateo-tavera/accounting-daily-tasks/entity/mocks"
	repo_mock "github.com/mateo-tavera/accounting-daily-tasks/repository/mocks"

	"github.com/mateo-tavera/accounting-daily-tasks/service"
	"github.com/stretchr/testify/assert"
)

// Initialize global variables

func TestNewService(t *testing.T) {
	// Mock elements from the service
	repoMock := new(repo_mock.RepositoryService)

	// Execute function
	svcTest := service.NewService(repoMock)

	// Evaluate results
	assert.NotNil(t, svcTest)

}

func TestGetTask(t *testing.T) {

	t.Run("error 'Bad Request' - GetTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetTask", 1, "tech1").Return(nil, errors.New("error executing query"))
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.GetTask(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 400)
		assert.Nil(t, err)

	})

	t.Run("error 'Unauthorized - GetTask'", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetTask", 1, "tech1").Return(nil, errors.New("permission denied"))
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.GetTask(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 401)
		assert.Nil(t, err)

	})

	t.Run("error 'Request successfull - GetTask'", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetTask", 1, "tech1").Return(entity_mocks.GenerateSvcTaskOK(), nil)
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.GetTask(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 200)
		assert.Nil(t, err)

	})
}

func TestGetAllTasks(t *testing.T) {
	t.Run("error 'No tasks available' - GetAllTasks", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetAllTasks", 1, "tech1").Return(nil, nil)
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.GetAllTasks(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 200)
		assert.Nil(t, err)

	})

	t.Run("error 'Request successfull - GetTask'", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetAllTasks", 1, "tech1").Return(entity_mocks.GenerateSvcAllTasksOK(), nil)
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.GetAllTasks(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 200)
		assert.Nil(t, err)

	})
}

func TestCreateTask(t *testing.T) {

	t.Run("error 'Bad Request' - CreateTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("CreateTask", entity.Task{}).Return(0, errors.New("permission denied"))
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.CreateTask(entity.Task{})

		// Evaluate results
		assert.Equal(t, response.Code, 400)
		assert.Nil(t, err)

	})

	t.Run("error 'Task Created' - CreateTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("CreateTask", entity.Task{}).Return(1, nil)
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.CreateTask(entity.Task{})

		// Evaluate results
		assert.Equal(t, response.Code, 201)
		assert.Nil(t, err)

	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("error 'Cannot get task from db' - UpdateTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetTask", 1, "tech1").Return(nil, errors.New("error executing query"))
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.UpdateTask(entity.Task{Id: 1, User: "tech1"})

		// Evaluate results
		assert.Equal(t, response.Code, 400)
		assert.Nil(t, err)

	})

	t.Run("error 'Bad Request to db' - UpdateTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetTask", 1, "tech1").Return(entity_mocks.GenerateSvcTaskOK(), nil)
		repoMock.On("UpdateTask", entity.Task{Id: 1, User: "tech1"}).Return(errors.New("error executing query"))
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.UpdateTask(entity.Task{Id: 1, User: "tech1"})

		// Evaluate results
		assert.Equal(t, response.Code, 400)
		assert.Nil(t, err)

	})

	t.Run("error 'Updated Task succesfully' - UpdateTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("GetTask", 1, "tech1").Return(entity_mocks.GenerateSvcTaskOK(), nil)
		repoMock.On("UpdateTask", entity.Task{Id: 1, User: "tech1"}).Return(nil)
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.UpdateTask(entity.Task{Id: 1, User: "tech1"})

		// Evaluate results
		assert.Equal(t, response.Code, 200)
		assert.Nil(t, err)

	})

}

func TestDeleteTask(t *testing.T) {
	t.Run("error 'Bad Request' - DeleteTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("DeleteTask", 1, "tech1").Return(errors.New("error executing query"))
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.DeleteTask(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 400)
		assert.Nil(t, err)

	})

	t.Run("error 'Task deleted succesfully' - DeleteTask", func(t *testing.T) {
		// Mock elements for the service
		repoMock := new(repo_mock.RepositoryService)
		repoMock.On("DeleteTask", 1, "tech1").Return(nil)
		svcTest := service.NewService(repoMock)

		// Execute function
		response, err := svcTest.DeleteTask(1, "tech1")

		// Evaluate results
		assert.Equal(t, response.Code, 200)
		assert.Nil(t, err)

	})
}
