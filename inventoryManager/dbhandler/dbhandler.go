package dbhandler

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const SELECT_QUERY_TIMEOUT_SECS = 5
const RESERVATION_TRANSACTION_TIMEOUT_SECS = 10
const REMOVE_RESERVATION_TIMEOUT_SECS = 5

func InitDBHandler(pool *sql.DB) DBHandler {
	db := DBHandler{dbPool: pool}

	return db
}

type DBHandler struct {
	dbPool *sql.DB
}

func (dbHandler DBHandler) query(cntx context.Context, productID string) (int32, error) {
	ctx, cancel := context.WithTimeout(cntx, SELECT_QUERY_TIMEOUT_SECS*time.Second)
	defer cancel()

	var quantity int32
	row := dbHandler.dbPool.QueryRowContext(ctx, fmt.Sprintf(`SELECT quantity FROM product_details WHERE product_id = '%s'`, productID))
	err := row.Scan(&quantity)

	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("No items found for product %s", productID)
	}
	return quantity, err
}

func (dbHandler DBHandler) GetQuantityFromDB(cntx context.Context, productID string) (int32, error) {
	var err error
	var quantity int32

	quantity, err = dbHandler.query(cntx, productID)
	return quantity, err
}

func (dbHandler DBHandler) GetReservationToken(cntx context.Context, productID string, quantity int32) (string, error) {
	ctx, cancel := context.WithTimeout(cntx, RESERVATION_TRANSACTION_TIMEOUT_SECS*time.Second)
	defer cancel()

	tx, err := dbHandler.dbPool.BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Failed to start transaction %s", err))
	}

	var inventoryQuantity int
	row := dbHandler.dbPool.QueryRowContext(ctx, fmt.Sprintf(`SELECT quantity FROM product_details WHERE product_id = '%s'`, productID))
	err = row.Scan(&inventoryQuantity)
	if err == sql.ErrNoRows || inventoryQuantity == 0 {
		tx.Rollback()
		return "", fmt.Errorf("No items available for product %s", productID)
	}

	query := fmt.Sprintf(`select sum(quantity) from reserve_purchase where product_id=%s`, productID)
	row = tx.QueryRowContext(cntx, query)
	var reserved int

	err = row.Scan(&reserved)
	if err == sql.ErrNoRows {
		reserved = 0
	}

	availableQuantity := inventoryQuantity - reserved
	if int32(availableQuantity) < quantity {
		tx.Rollback()
		return "", fmt.Errorf("Not enough items available for product %s", productID)
	}

	var reservationID = uuid.New().String()

	var stmt *sql.Stmt
	stmt, err = tx.Prepare(`INSERT INTO reserve_purchase set reserve_id=?, product_id=?, quantity=?`)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	_, err = stmt.Exec(reservationID, productID, quantity)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return reservationID, nil
}

func (dbHandler DBHandler) CommitReservation(cntx context.Context, reservationID string) error {
	ctx, cancel := context.WithTimeout(cntx, REMOVE_RESERVATION_TIMEOUT_SECS*time.Second)
	defer cancel()

	var err error

	tx, err := dbHandler.dbPool.BeginTx(cntx, nil)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to start transaction %s", err))
	}

	var productID string
	var quantity int = 1

	rows := tx.QueryRowContext(ctx, fmt.Sprintf(`SELECT product_id, quantity FROM reserve_purchase WHERE reserve_id="%s"`, reservationID))
	err = rows.Scan(&productID, &quantity)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to get reservation details: %s", err)
	}

	var inventoryQuantity int
	row := dbHandler.dbPool.QueryRowContext(ctx, fmt.Sprintf(`SELECT quantity FROM product_details WHERE product_id = '%s'`, productID))
	err = row.Scan(&inventoryQuantity)

	var stmt *sql.Stmt
	stmt, err = tx.Prepare(`UPDATE product_details SET quantity = ? WHERE product_id = ?`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to update product details: %s", err)
	}

	_, err = stmt.Exec(inventoryQuantity-quantity, productID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to update product details: %s", err)
	}

	stmt, err = tx.Prepare(`DELETE FROM reserve_purchase WHERE reserve_id = ?`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to remove reservation: %s", err)
	}

	_, err = stmt.Exec(reservationID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to remove reservation: %s", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to commit reservation: %s", err)
	}

	return nil
}

func (dbHandler DBHandler) RemoveReservation(cntx context.Context, reservationID string) error {
	ctx, cancel := context.WithTimeout(cntx, REMOVE_RESERVATION_TIMEOUT_SECS*time.Second)
	defer cancel()
	var err error

	tx, err := dbHandler.dbPool.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to start transaction %s", err))
	}

	var stmt *sql.Stmt
	stmt, err = tx.Prepare(`DELETE FROM reserve_purchase WHERE RESERVE_ID = ?`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf(fmt.Sprintf("Failed to prepare statement: %s", err))
	}

	_, err = stmt.Exec(reservationID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf(fmt.Sprintf("Failed to executed statement: %s", err))
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to rollback transaction: %s", err))
	}

	return nil
}
