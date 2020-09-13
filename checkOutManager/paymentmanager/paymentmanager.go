package paymentmanager

import (
	"AngelProj/checkOutManager/grpcinventoryclient"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type Item struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CheckOutJob struct {
	Items []Item `json:"items"`
}

type CheckOutResponse struct {
	Done bool   `json:"done"`
	Err  string `json:"error"`
}

func grpcGetToken(productID string, quantity int32) (string, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial("inventoryManager:9100", grpc.WithInsecure())
	if err != nil {
		return "", err
	}

	defer conn.Close()

	c := grpcinventoryclient.NewInventoryServiceClient(conn)
	request := grpcinventoryclient.ReservationRequest{ProductID: productID, Quantity: quantity}
	response, err := c.GetResevationToken(context.Background(), &request)
	if err != nil {
		return "", err
	}

	return response.ReservationID, err
}

func CommitReservation(reservationToken string) (bool, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial("inventoryManager:9100", grpc.WithInsecure())
	if err != nil {
		return false, err
	}

	defer conn.Close()

	c := grpcinventoryclient.NewInventoryServiceClient(conn)
	request := grpcinventoryclient.CommitReservationRequest{ReservationID: reservationToken}
	_, err = c.CommitReservation(context.Background(), &request)
	if err != nil {
		return false, err
	}

	return true, err
}

func RollBackReservation(reservationToken string) (bool, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial("inventoryManager:9100", grpc.WithInsecure())
	if err != nil {
		return false, err
	}

	defer conn.Close()

	c := grpcinventoryclient.NewInventoryServiceClient(conn)
	request := grpcinventoryclient.RollBackReservationRequest{ReservationID: reservationToken}
	_, err = c.RollBackReservation(context.Background(), &request)
	if err != nil {
		return false, err
	}

	return true, err
}

func payProcess() error {
	time.Sleep(2 * time.Second)

	return nil
}

func ProcessPayment(job CheckOutJob) error {
	var err error
	var reservationTokens []string

	for _, item := range job.Items {
		fmt.Println("PID::", item.ProductID, item.Quantity)
		reservationToken, e := grpcGetToken(item.ProductID, int32(item.Quantity))
		if e != nil {
			fmt.Println("ERR::", e)
			return e
		}

		reservationTokens = append(reservationTokens, reservationToken)
	}

	fmt.Println(reservationTokens)
	err = payProcess()
	fmt.Println(err)

	for _, token := range reservationTokens {
		fmt.Println("TK::", token)
		if err != nil {
			RollBackReservation(token)
			return err
		}
		success, e := CommitReservation(token)
		fmt.Println("S::", success)
		if e != nil {
			err = e
		}
	}

	fmt.Println("EE::", err)
	return err
}
