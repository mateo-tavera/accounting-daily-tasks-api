package endpoint

import (
	"context"
	"fmt"

	"github.com/mateo-tavera/accounting-daily-tasks/entity"
	service "github.com/mateo-tavera/accounting-daily-tasks/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateTaskEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Task)
		result, err := svc.CreateTask(req)
		if err != nil {
			return result, fmt.Errorf("service failed: %v", err)
		}
		return &entity.Response{Message: result.Message, Code: result.Code, Data: result.Data}, nil
	}

}

func MakeGetTaskEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Task)
		result, err := svc.GetTask(req.Id, req.User)
		if err != nil {
			return result, fmt.Errorf("service failed: %v", err)
		}
		return &entity.Response{Message: result.Message, Code: result.Code, Data: result.Data}, nil
	}
}

func MakeGetAllTasksEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Task)
		result, err := svc.GetAllTasks(req.Id, req.User)
		if err != nil {
			return nil, fmt.Errorf("service failed: %v", err)
		}
		return &entity.Response{Message: result.Message, Code: result.Code, Data: result.Data}, nil
	}
}

func MakeUpdateTaskEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Task)
		result, err := svc.UpdateTask(req)
		if err != nil {
			return nil, fmt.Errorf("service failed: %v", err)
		}
		return &entity.Response{Message: result.Message, Code: result.Code}, nil
	}
}

func MakeDeleteTaskEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Task)
		result, err := svc.DeleteTask(req.Id, req.User)
		if err != nil {
			return nil, fmt.Errorf("service failed: %v", err)
		}
		return &entity.Response{Message: result.Message, Code: result.Code}, nil
	}
}
