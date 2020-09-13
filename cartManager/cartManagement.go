package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"

	"AngelProj/addcart/jobhandler"
	"AngelProj/addcart/joblistener"

	_ "github.com/go-sql-driver/mysql"
)

// var dbHost string = os.Getenv("DBHOST")
// var dbUserID string = os.Getenv("DBUSERID")
// var dbPassword = os.Getenv("DBPASSWORD")
// var dbName = os.Getebd("DBNAME")
// var parallelJobsToProcess = os.Getenv("CONCURRENCY")

var dbHost string = "mysql-development"
var dbUserID string = "root"
var dbPassword = "pass"
var dbName = "testapp"
var parallelJobsToProcess = "1"

func setDBPoolConfig(pool *sql.DB, concurrencyCount int) {
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(concurrencyCount)
	pool.SetMaxOpenConns(concurrencyCount)
}

func getDBConnectionPool() (*sql.DB, error) {
	dbPool, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql-development:3306)/%s", dbUserID, dbPassword, dbName))
	if err != nil {
		fmt.Println("DB conn error ... cartmanagement ", err)
	}
	return dbPool, err
}

func failOnError(err error, message string) {
	if err != nil {
		log.Printf("%s : %v", message, err)
	}
}

func panicOnError(err error, message string) {
	if err != nil {
		log.Fatal("%s : %v", message, err)
	}
}

func getRMQConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	return conn, err
}

func setRMQChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	return ch, err
}

func setRMQQueue(ch *amqp.Channel) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"cart_manager_queue", // name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)

	return q, err
}

func setRMQQueuePrefetchCount(ch *amqp.Channel, prefetchCount int) error {
	err := ch.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size
		false,         // global
	)
	return err
}

func registerConsumer(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	return msgs, err
}

/*
	Cart Management is a scalable worker which adds an item to user cart
	Validations:
		1. The available quantity needs to be > 0
		2. If an item already exists in cart, item quantity is incremented by 1
*/
func main() {
	var err error

	var concurrencyCount int
	concurrencyCount, err = strconv.Atoi(parallelJobsToProcess)
	if err != nil {
		concurrencyCount = 1
	}

	var rmqConnection *amqp.Connection
	rmqConnection, err = getRMQConnection()
	panicOnError(err, "Failed to establish connection with RMQ")
	defer rmqConnection.Close()

	var rmqChannel *amqp.Channel
	rmqChannel, err = setRMQChannel(rmqConnection)
	panicOnError(err, "Failed to set RMQ Channel")
	defer rmqChannel.Close()

	var rmqQueue amqp.Queue
	rmqQueue, err = setRMQQueue(rmqChannel)
	panicOnError(err, "Failed to set RMQ Queue")

	// Prefetch sets the number of jobs RMQ will assign maximum to a worker
	// The worker here can process multiple jobs as per the concurrency count
	err = setRMQQueuePrefetchCount(rmqChannel, concurrencyCount)
	panicOnError(err, "Failed to set RMQ Queue PrefetchCount")

	var messageChannel <-chan amqp.Delivery
	messageChannel, err = registerConsumer(rmqChannel, rmqQueue)
	failOnError(err, "Failed to register consumer")

	// err = rmqChannel.ExchangeDeclare(
	// 	"cart-manager-results", // exchange-name
	// 	"fanout",               //exchange type
	// 	true,                   // durable
	// 	false,                  // auto-delete
	// 	false,                  // internal
	// 	false,                  // no-wait
	// 	nil,                    // arguments
	// )
	// panicOnError(err, "Failed to set RMQ Result Exchange")

	var dbPool *sql.DB
	dbPool, err = getDBConnectionPool()
	panicOnError(err, "Failed to connect to DB")
	defer dbPool.Close()

	setDBPoolConfig(dbPool, concurrencyCount)

	jobHandler := jobhandler.JobHandler{RmqChannel: rmqChannel}
	var inputChannel chan joblistener.Job
	var outputChannel chan joblistener.JobResponse
	inputChannel, outputChannel = jobHandler.SetupWorkers(dbPool, concurrencyCount)

	defer close(inputChannel)
	defer close(outputChannel)

	var forEver chan bool

	go joblistener.ListenAndProcessJobs(messageChannel, inputChannel)
	for response := range outputChannel {
		fmt.Println(response)
		err = rmqChannel.Publish(
			"",
			"cart_result_queue",
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/json",
				CorrelationId: response.RequestID,
				Body:          []byte(fmt.Sprintf("%s", response)),
			},
		)

		response.AmqpMessage.Ack(false)
	}

	log.Println("Waiting for jobs.... Press [Cntrl]C to exit")
	// Stuck for ever till some one forces the code to exit
	<-forEver
}
