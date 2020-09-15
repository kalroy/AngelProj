package dbhandler

import (
	"database/sql"
	"fmt"
)

func startTransaction(conn *sql.DB) (*sql.Tx, error) {
	return conn.Begin()
}

func AddToCart(conn *sql.DB, clientID string, productID string, quantity int) error {
	var err error
	var transaction *sql.Tx
	transaction, err = startTransaction(conn)
	if err != nil {
		return err
	}

	var stmt *sql.Stmt
	stmt, err = transaction.Prepare(`INSERT INTO client_cart set client_id=?, product_id=?, quantity=? ON DUPLICATE KEY UPDATE quantity = quantity + ?`)
	if err != nil {
		transaction.Rollback()
		return err
	}
	_, err = stmt.Exec(clientID, productID, quantity, quantity)
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return err
	}
	return err
}

func GetCartDetails(conn *sql.DB, clientID string) (map[string]int, error) {
	var cartDetails map[string]int = make(map[string]int)

	rows, err := conn.Query(fmt.Sprintf(`SELECT product_id, quantity FROM client_cart WHERE client_id = '%s'`, clientID))
	if err != nil {
		fmt.Printf("Failed to query: %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var productID string
		var quantity int

		err = rows.Scan(&productID, &quantity)
		if err != nil {
			continue
		}

		fmt.Println(productID, quantity)
		cartDetails[productID] = quantity
	}

	return cartDetails, nil
}

func RemoveCart(conn *sql.DB, clientID string) error {
	var err error
	var transaction *sql.Tx
	transaction, err = startTransaction(conn)
	if err != nil {
		return err
	}

	var stmt *sql.Stmt
	stmt, err = transaction.Prepare(`DELETE FROM client_cart WHERE client_id=?`)
	if err != nil {
		transaction.Rollback()
		return err
	}
	_, err = stmt.Exec(clientID)
	if err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return err
	}
	return err
}
