package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	ep "github.com/mateo-tavera/accounting-daily-tasks/endpoint"
	e "github.com/mateo-tavera/accounting-daily-tasks/entity"

	"github.com/gorilla/mux"
	"github.com/mateo-tavera/accounting-daily-tasks/service"
)

func DecodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := castRequest(r)
	request.Date = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %v", err)
	}

	return request, nil
}

func DecodeGetTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := castRequest(r)
	id := mux.Vars(r)["id"]
	request.Id, _ = strconv.Atoi(id)

	return request, nil
}

func DecodeUpdateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := castRequest(r)
	id := mux.Vars(r)["id"]
	request.Id, _ = strconv.Atoi(id)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request: %v", err)
	}

	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	resp := response.(*e.Response)
	w.WriteHeader(resp.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}

func MakeHandler(svc service.Service) http.Handler {

	createTaskHandler := httptransport.NewServer(
		ep.MakeCreateTaskEndpoint(svc),
		DecodeCreateTaskRequest,
		EncodeResponse,
	)

	getTaskHandler := httptransport.NewServer(
		ep.MakeGetTaskEndpoint(svc),
		DecodeGetTaskRequest,
		EncodeResponse,
	)

	getAllTasksHandler := httptransport.NewServer(
		ep.MakeGetAllTasksEndpoint(svc),
		DecodeGetTaskRequest,
		EncodeResponse,
	)

	updateTaskHandler := httptransport.NewServer(
		ep.MakeUpdateTaskEndpoint(svc),
		DecodeUpdateTaskRequest,
		EncodeResponse,
	)

	deleteTaskHandler := httptransport.NewServer(
		ep.MakeDeleteTaskEndpoint(svc),
		DecodeGetTaskRequest,
		EncodeResponse,
	)

	// Create main route
	r := mux.NewRouter()

	// Create subroutes depending on authorization level
	// managers := r.NewRoute().Subrouter()
	// technicians := r.NewRoute().Subrouter()

	// Set middleware
	// managers.Use(service.AuthorizationMgrMiddleware)     // Only the routes for 'managers' use this middleware
	// technicians.Use(service.AuthorizationTechMiddleware) // Only the routes for 'technicians' use this middleware

	// API endpoints
	r.Methods(http.MethodPost).Path("/task").Handler(createTaskHandler)
	r.Methods(http.MethodGet).Path("/task/{id}").Handler(getTaskHandler)
	r.Methods(http.MethodGet).Path("/task").Handler(getAllTasksHandler)
	r.Methods(http.MethodPatch).Path("/task/{id}").Handler(updateTaskHandler)
	r.Methods(http.MethodDelete).Path("/task/{id}").Handler(deleteTaskHandler)

	return r
}

// Cast request as type Task. Used to save code
func castRequest(r *http.Request) e.Task {
	user := r.Header.Get("Authorization")
	var request e.Task
	request.User = user
	return request
}
