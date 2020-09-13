package jobhandler

import (
	"AngelProj/addcart/dbhandler"
	"AngelProj/addcart/joblistener"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type item struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type checkOutJob struct {
	Items []item `json:"items"`
}

type checkOutJobResponse struct {
	Done bool   `json:"done`
	Err  string `json:"error"`
}

func checkOutItems(conn *sql.DB, channel *amqp.Channel, results <-chan amqp.Delivery, queue amqp.Queue, job joblistener.Job) error {
	cartDetails, err := dbhandler.GetCartDetails(conn, job.ClientID)
	if err != nil {
		return err
	}

	var j = checkOutJob{}

	var items []item
	for productID, quantity := range cartDetails {
		items = append(items, item{ProductID: productID, Quantity: quantity})
	}

	j.Items = items

	checkOutJob, _ := json.Marshal(j)
	channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			CorrelationId: job.RequestID,
			ContentType:   "text/plain",
			Body:          checkOutJob,
		},
	)

	fmt.Println("Reading results..")
	var done bool = true
	var e string
	for result := range results {
		if job.RequestID == result.CorrelationId {
			var res checkOutJobResponse = checkOutJobResponse{}
			json.Unmarshal(result.Body, &res)
			done = res.Done
			e = res.Err
			result.Ack(false)
			break
		}
	}

	fmt.Println("D:: ", done)
	fmt.Println("E:: ", e)
	if !done {
		return fmt.Errorf(e)
	}
	return nil
}
