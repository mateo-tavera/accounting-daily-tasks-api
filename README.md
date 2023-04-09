# Accounting Daily Tasks API

The Accounting Daily Tasks API is a Go application that connects to a MySQL database and provides endpoints to manage daily accounting tasks. This API allows CRUD (Create, Read, Update, Delete) operations on a MySQL database to keep track of the tasks, according to the user (manager or technician).
Within the code, techincians are represented with `tech` and managers with `mgr`

## Implementation

To run the Accounting Daily Tasks API, you will need the following requirements:

- Run the docker compose file entering this command in a terminal `docker compose up -d`. This includes the container for the app and for the data base, with the volume for persistence.

### Go dependencies:
The API uses some Go dependencies, such as:.
- `github.com/go-kit/kit` and `github.com/gorilla/mux`: for microservice architecture and routing
- `github.com/go-sql-driver/mysql`: to interact with the MySQL database.
- `github.com/DATA-DOG/go-sqlmock` and `github.com/vektra/mockery`: for mocking the interfaces and use sql mock methods

## Endpoints

Before running the Accounting Daily Tasks API, you need to know how it works, so here is the explaination of how to use it, according to the endpoints
Fisrts, the security and authentication works using header parameters. This means that in order to get access to the endpoints, you must register your user in the header `Authorization`. This will be decoded into the sturct `Task`

```go
type Task struct {
	Id      int       `json:"id,omitempty"`
	User    string    `json:"user" validate:"required"`
	Summary string    `json:"summary" validate:"required"` // TODO: encrypt
	Date    time.Time `json:"date"`
	Status  string    `json:"status" validate:"required"`
}
```

### 1. Create task: 
#### /task (POST)

This operation allows a technician to create a task for itself, providing the necessary attribuites in JSON format. It is represented in the next example:
```json
{
    "summary":"fix bugs",
    "status":"pending"
}
```
Note that the User, Id and Date attributes are missing in the JSON request. That is because Id is generated automatically by the database; the date is registered once the request is done; and the user is provided in the headers, as mentioned before. This only works for technicians registered.
_Try it with other users to check error validation_

Response:
![image](https://user-images.githubusercontent.com/92878710/230799400-d9e59cbf-1f0b-48ad-9a83-ceacfe3c8872.png)


### 2. Get task: 
#### /task/{id} (GET)
This method recovers the inforamtion of a task for a given Id. This works both for technicians and managers, but technicians can only see their taks, while managers can see any task. It uses no body for the request, but the authentication is validated with header `Authorization`.
_Try it with other users or differente Ids to check error validation_

Response:
![image](https://user-images.githubusercontent.com/92878710/230799415-ff1b6cf4-7d14-47f3-be94-bd98bd77ab6f.png)



### 3. Get All tasks: 
#### /task (GET)
This feature only works for managers and it recovers all tasks registered in the database. It returns a list of tasks to any manager register in the authorization list
_Try it with other users or differente Ids to check error validation_

Response:
![image](https://user-images.githubusercontent.com/92878710/230799479-099bbcbc-71ce-4fdb-925f-f9917f3eb84e.png)


### 4. Update task: 
#### /task/{id} (PATCH)
A task can only be updated by the technician who is assigned, so it must be regsitered in the authorization list and being declare in the header `Authorization`. The only values that can be modifed are the Summary and the Status. Date is updated automatically and Id is immutable.
The request must have JSON format specifying only the atributes to be updated.
```json
{
    "status":"done"
}
```
_Try it with other users or differente Ids to check error validation_
Response:
![image](https://user-images.githubusercontent.com/92878710/230799563-749d8a33-52ce-476b-8e1c-8af4ea20d25e.png)

### 5. Delete task: 
#### /task/{id} (DELETE)
This feature can be only performed by managers and is required the authentication and the Id of the task to be deleted.
_Try it with other users or differente Ids to check error validation_
Response:
![image](https://user-images.githubusercontent.com/92878710/230799621-36121fc8-330d-4c3f-9b2e-10d0c72b0b7e.png)

## TODOs:
To improve performance and quality of the code, there are some aspects that can be changed in order to provide more security:
- The Summary attribute can be encrypted to protect the integrity of the task, in case of having personal information. This could be done either by the database, declaring column as VARBINARY or BLOB; or using the encryption library integrated in Go.
- The users (managers and techinicians) could be authenticated by using Access Tokens, and being stored in a different encrypted database.
- Use GORM to connect with database and perfom queris.

