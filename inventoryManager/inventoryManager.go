package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"AngelProj/inventoryManager/dbhandler"
	"AngelProj/inventoryManager/grpcinventoryserver"
)

var dbHost string = "mysql-development"
var dbUserID string = "root"
var dbPassword = "pass"
var dbName = "testapp"
var parallelJobsToProcess = "3"

func setDBPoolConfig(pool *sql.DB, concurrencyCount int) {
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(concurrencyCount)
	pool.SetMaxOpenConns(concurrencyCount)
}

func getDBConnectionPool() (*sql.DB, error) {
	dbPool, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql-development:3306)/%s", dbUserID, dbPassword, dbName))
	if err != nil {
		fmt.Println("database connection error ... inventory ", err)
	}
	return dbPool, err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s : %s", msg, err))
	}
}

func panicOnError(err error, message string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s : %s", message, err))
	}
}

func main() {
	var err error

	var dbPool *sql.DB
	dbPool, err = getDBConnectionPool()
	panicOnError(err, "Failed to connect to DB")
	defer dbPool.Close()

	dbHandler := dbhandler.InitDBHandler(dbPool)

	lis, err := net.Listen("tcp", ":9100")
	failOnError(err, "Failed to bind to the listening port 9100")

	inv := grpcinventoryserver.Server{DBHandler: dbHandler}

	server := grpc.NewServer()
	grpcinventoryserver.RegisterInventoryServiceServer(server, &inv)

	err = server.Serve(lis)
	failOnError(err, "Failed to start grpc server at port 9100")
}
