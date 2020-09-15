package grpcinventoryserver

import (
	"AngelProj/inventoryManager/dbhandler"
	"context"
)

type Server struct {
	DBHandler dbhandler.DBHandler
}

func (s *Server) GetAvailableQuantity(cntx context.Context, request *QuantityRequest) (*QuantityResponse, error) {
	var response *QuantityResponse = &QuantityResponse{}

	var productID string = request.GetProductID()
	quantity, err := s.DBHandler.GetQuantityFromDB(cntx, productID)
	if err != nil {
		return response, err
	}
	response.ProductID = productID
	response.Quantity = quantity

	return response, nil
}

func (s *Server) GetResevationToken(cntx context.Context, request *ReservationRequest) (*ReservationResponse, error) {
	var response *ReservationResponse = &ReservationResponse{}

	var productID string = request.GetProductID()
	var quantity int32 = request.GetQuantity()

	var token string
	var err error
	token, err = s.DBHandler.GetReservationToken(cntx, productID, quantity)
	if err != nil {
		return response, err
	}

	response.ReservationID = token
	return response, nil
}

func (s *Server) RollBackReservation(cntx context.Context, request *RollBackReservationRequest) (*RollBackReservationResponse, error) {
	var response *RollBackReservationResponse = &RollBackReservationResponse{Success: true}

	var reservationToken string = request.GetReservationID()

	err := s.DBHandler.RemoveReservation(cntx, reservationToken)

	if err != nil {
		response.Success = false
	}
	return response, err
}

func (s *Server) CommitReservation(cntx context.Context, request *CommitReservationRequest) (*CommitReservationResponse, error) {
	var response *CommitReservationResponse = &CommitReservationResponse{Success: true}

	var reservationToken string = request.GetReservationID()

	err := s.DBHandler.CommitReservation(cntx, reservationToken)
	if err != nil {
		response.Success = false
	}

	return response, err
}
