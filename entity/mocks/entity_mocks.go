package mocks

import (
	"net/http"
	"time"

	"github.com/mateo-tavera/accounting-daily-tasks/entity"
)

func GenerateSvcResponsetOK() *entity.Response {
	return &entity.Response{
		Code:    http.StatusOK,
		Message: "Request has been completed successfully",
		Data:    nil,
	}
}

func GenerateSvcResponsetNOK() *entity.Response {
	return &entity.Response{
		Code:    http.StatusBadRequest,
		Message: "error executing query",
		Data:    nil,
	}
}

func GenerateSvcTaskOK() *entity.Task {
	return &entity.Task{
		Id:      1,
		User:    "tech1",
		Summary: "summary",
		Date:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		Status:  "status",
	}
}

func GenerateCreateTaskOK() entity.Task {
	return entity.Task{
		User:    "tech1",
		Summary: "summary",
		Date:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		Status:  "status",
	}
}

func GenerateCreateTaskNOK() entity.Task {
	return entity.Task{
		User:    "tech11",
		Summary: "summary",
		Date:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		Status:  "status",
	}
}

func GenerateSvcAllTasksOK() []entity.Task {
	return []entity.Task{{
		Id:      1,
		User:    "tech1",
		Summary: "summary",
		Date:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		Status:  "status",
	}, {
		Id:      2,
		User:    "tech1",
		Summary: "summary",
		Date:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		Status:  "status",
	}}
}
