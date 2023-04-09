package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mateo-tavera/accounting-daily-tasks/entity"

	repo "github.com/mateo-tavera/accounting-daily-tasks/repository"
)

type Service interface {
	GetTask(int, string) (*entity.Response, error)
	GetAllTasks(int, string) (*entity.Response, error)
	CreateTask(entity.Task) (*entity.Response, error)
	UpdateTask(entity.Task) (*entity.Response, error)
	DeleteTask(int, string) (*entity.Response, error)
}

type service struct {
	repository repo.RepositoryService
}

// Interface implementation
func NewService(repo repo.RepositoryService) Service {
	return &service{
		repository: repo,
	}
}

// Get single task based on id and user
func (svc *service) GetTask(id int, user string) (*entity.Response, error) {

	//Call the repository service
	task, err := svc.repository.GetTask(id, user)

	return svc.ValidateResponseGet(err, task)
}

// Get all tasks according to user
func (svc *service) GetAllTasks(id int, user string) (*entity.Response, error) {

	//Call the repository service
	result, err := svc.repository.GetAllTasks(id, user)

	// Validate if there are not tasks available
	if err == nil && result == nil {
		return &entity.Response{Code: http.StatusOK, Message: fmt.Sprintf("No tasks related to user %s", user), Data: result}, nil
	}

	return svc.ValidateResponseGet(err, result)
}

// Save a new task
func (svc *service) CreateTask(task entity.Task) (*entity.Response, error) {

	// Call respository service
	result, err := svc.repository.CreateTask(task)
	if err != nil {
		response := entity.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    result,
		}
		return &response, nil
	}

	return &entity.Response{Code: http.StatusCreated, Message: "Task Created", Data: result}, nil
}

// Update a task using PATCH
func (svc *service) UpdateTask(task entity.Task) (*entity.Response, error) {
	//Call the repository service
	result, _ := svc.GetTask(task.Id, task.User) // error evaluated in GetTask method
	if result.Code != 200 {
		// Concatenate the error got from GetTask function
		err := fmt.Errorf("error getting task: %v", result.Message)
		return svc.ValidateResponseGet(err, result)
	}

	//Update task in db
	err := svc.repository.UpdateTask(task)
	if err != nil {
		response := entity.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return &response, nil
	}

	return &entity.Response{Code: http.StatusOK, Message: "Task updated succesfully", Data: nil}, nil
}

// Delete a task
func (svc *service) DeleteTask(id int, user string) (*entity.Response, error) {

	//Call the repository service
	err := svc.repository.DeleteTask(id, user)
	if err != nil {
		response := entity.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return &response, nil
	}

	return &entity.Response{Code: http.StatusOK, Message: "Task deleted succesfully", Data: nil}, nil

}

func (svc *service) ValidateResponseGet(err error, task interface{}) (*entity.Response, error) {
	if err != nil {
		//Validate authorization
		if strings.Contains(err.Error(), "permission denied") {
			response := entity.Response{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			}
			return &response, nil
		} else {
			response := entity.Response{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
			return &response, nil
		}
	}
	return &entity.Response{Code: http.StatusOK, Message: "Request has been completed successfully", Data: task}, nil
}
