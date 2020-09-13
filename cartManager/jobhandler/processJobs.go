package jobhandler

import (
	"AngelProj/addcart/dbhandler"
	"AngelProj/addcart/joblistener"
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc"

	"AngelProj/addcart/grpcinventoryclient"

	"github.com/streadway/amqp"
)

type JobType string

const (
	AddToCart JobType = "addToCart"
	CheckOut          = "checkOut"
)

type JobHandler struct {
	RmqChannel *amqp.Channel
}

var checkoutQueue amqp.Queue
var checkoutResultQueue <-chan amqp.Delivery

func (handler JobHandler) grpcGetInventory(productID string) (int32, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial("inventoryManager:9100", grpc.WithInsecure())
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	c := grpcinventoryclient.NewInventoryServiceClient(conn)
	request := grpcinventoryclient.QuantityRequest{ProductID: productID}
	response, err := c.GetAvailableQuantity(context.Background(), &request)
	if err != nil {
		return 0, err
	}

	return response.Quantity, err
}

func (handler JobHandler) shouldAddItemToCart(productID string, quantity int) (bool, error) {
	qty, err := handler.grpcGetInventory(productID)
	return (err == nil && qty >= int32(quantity)), err
}

func (handler JobHandler) handleAddToCart(conn *sql.DB, job joblistener.Job, output chan joblistener.JobResponse) {
	var err error
	var shouldAdd bool
	shouldAdd, err = handler.shouldAddItemToCart(job.ProductID, job.Quantity)
	if shouldAdd {
		err = dbhandler.AddToCart(conn, job.ClientID, job.ProductID, job.Quantity)
	}

	var done bool = true
	var e string
	if err != nil {
		done = false
		e = err.Error()
	}

	response := joblistener.JobResponse{
		Done:        done,
		RequestID:   job.RequestID,
		Error:       e,
		AmqpMessage: job.AmqpMessage,
	}
	output <- response
}

func (handler JobHandler) handleCheckOut(conn *sql.DB, job joblistener.Job, output chan joblistener.JobResponse) {
	err := checkOutItems(conn, handler.RmqChannel, checkoutResultQueue, checkoutQueue, job)

	var done bool = true
	var errString string

	if err != nil {
		done = false
		errString = err.Error()
	}

	response := joblistener.JobResponse{
		Done:        done,
		RequestID:   job.RequestID,
		Error:       errString,
		AmqpMessage: job.AmqpMessage,
	}
	output <- response
}

func (handler JobHandler) handleUnknownJob(job joblistener.Job, output chan joblistener.JobResponse) {
	var done = false
	var err = fmt.Sprintf("Unknown Job Type %s", job.JobType)

	response := joblistener.JobResponse{
		Done:        done,
		RequestID:   job.RequestID,
		Error:       err,
		AmqpMessage: job.AmqpMessage,
	}
	output <- response
}

func (handler JobHandler) jobProcessor(conn *sql.DB,
	id int,
	input chan joblistener.Job,
	output chan joblistener.JobResponse) error {
	var err error

	checkoutQueue, err = handler.RmqChannel.QueueDeclare(
		"checkout_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to intiate Checkout Queue : %s", err.Error())
	}

	q, err := handler.RmqChannel.QueueDeclare(
		"checkout-results", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // noWait
		nil,                // arguments
	)

	checkoutResultQueue, err = handler.RmqChannel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for job := range input {
		switch JobType(job.JobType) {
		case AddToCart:
			handler.handleAddToCart(conn, job, output)
			break
		case CheckOut:
			handler.handleCheckOut(conn, job, output)
			break
		default:
			handler.handleUnknownJob(job, output)
		}
	}

	return nil
}

func (handler JobHandler) SetupWorkers(conn *sql.DB,
	concurrencyCount int) (chan joblistener.Job, chan joblistener.JobResponse) {
	var input chan joblistener.Job = make(chan joblistener.Job)
	var output chan joblistener.JobResponse = make(chan joblistener.JobResponse)

	for i := 1; i <= concurrencyCount; i++ {
		go handler.jobProcessor(conn, i, input, output)
	}

	return input, output
}
