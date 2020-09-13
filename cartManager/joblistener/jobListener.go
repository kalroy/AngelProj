package joblistener

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// Job is the structure to follow the request data
type Job struct {
	RequestID   string        `json:"requestid"`
	JobType     string        `json:"jobtype"`
	ClientID    string        `json:"clientid"`
	ProductID   string        `json:"productid"`
	Quantity    int           `json:"quantity"`
	AmqpMessage amqp.Delivery `json:"-"`
}

func (j Job) String() string {
	e, _ := json.Marshal(&j)
	return string(e)
}

type JobResponse struct {
	Done        bool          `json:"isSuccess"`
	RequestID   string        `json:"requestid"`
	Error       string        `json:"error"`
	AmqpMessage amqp.Delivery `json:"-"`
}

func (j JobResponse) String() string {
	e, _ := json.Marshal(&j)
	return string(e)
}

func ListenAndProcessJobs(messageChannel <-chan amqp.Delivery, inputChannel chan Job) {
	for m := range messageChannel {
		j := Job{}
		json.Unmarshal(m.Body, &j)
		j.AmqpMessage = m
		inputChannel <- j
	}
}
