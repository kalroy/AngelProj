package main

import (
	"encoding/json"
	"fmt"
	"log"

	"AngelProj/checkOutManager/paymentmanager"

	"github.com/streadway/amqp"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Printf("%s : %v", message, err)
	}
}

func panicOnError(err error, message string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s : %v", message, err))
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
		"checkout_queue", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
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

func main() {

	const concurrencyCount = 1
	var err error

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

	forever := make(chan bool)

	fmt.Println("About to start reading messages")
	go func() {
		for job := range messageChannel {
			j := paymentmanager.CheckOutJob{}
			e := json.Unmarshal(job.Body, &j)
			if e != nil {
				fmt.Println(e.Error())
			}
			err = paymentmanager.ProcessPayment(j)

			done := err == nil

			var errString string
			if err != nil {
				errString = err.Error()
			}

			r := paymentmanager.CheckOutResponse{Done: done, Err: errString}
			body, _ := json.Marshal(r)
			rmqChannel.Publish(
				"",                 // exchange
				"checkout-results", // routing key
				false,              // mandatory
				false,              // immediate
				amqp.Publishing{
					ContentType:   "text/json",
					CorrelationId: job.CorrelationId,
					Body:          body,
				},
			)
			job.Ack(false)
		}
	}()

	<-forever

}
