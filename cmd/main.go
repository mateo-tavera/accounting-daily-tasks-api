package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mateo-tavera/accounting-daily-tasks/handler"
	repo "github.com/mateo-tavera/accounting-daily-tasks/repository"
	"github.com/mateo-tavera/accounting-daily-tasks/service"
)

func main() {
	// Create db connection
	db, err := getDBConnection()
	if err != nil {
		log.Fatalf("cannot establish connection: %v", err)
	}
	defer db.Close()

	// Create service and repo instance
	repository := repo.NewRepository(db)
	svc := service.NewService(repository)

	// Resgiter the routes to the service
	handler := handler.MakeHandler(svc)

	// Start server
	log.Printf("Listening at port: %d", 8081)

	httpListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		os.Exit(1)
	}
	log.Fatal(http.Serve(httpListener, handler))

}

func getDBConnection() (*sql.DB, error) {

	// TODO: create config file
	var (
		host     = "mysql-db"
		port     = "3306"
		user     = "user"
		password = "password"
		dbname   = "db_dailytasks"
	)

	mysqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	return sql.Open("mysql", mysqlConnection)
}
