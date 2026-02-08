package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"simple-cashier-api/models"

	"github.com/lib/pq"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("product id %d not found", item.ProductID)
			}
			return nil, err
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	var createdAt time.Time
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id, created_at", totalAmount).Scan(&transactionID, &createdAt)
	if err != nil {
		return nil, err
	}

	txIDs := make([]int, len(details))
	productIDs := make([]int, len(details))
	quantities := make([]int, len(details))
	subtotals := make([]int, len(details))

	for i, d := range details {
		txIDs[i] = transactionID
		productIDs[i] = d.ProductID
		quantities[i] = d.Quantity
		subtotals[i] = d.Subtotal
	}

	rows, err := tx.Query(
		`INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal)
						SELECT * FROM unnest($1::int[], $2::int[], $3::int[], $4::int[])
						RETURNING id, transaction_id, product_id, quantity, subtotal`,
		pq.Array(txIDs), pq.Array(productIDs), pq.Array(quantities), pq.Array(subtotals))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var insertedDetails []models.TransactionDetail
	for rows.Next() {
		var d models.TransactionDetail
		err = rows.Scan(&d.ID, &d.TransactionID, &d.ProductID, &d.Quantity, &d.Subtotal)
		if err != nil {
			return nil, err
		}

		insertedDetails = append(insertedDetails, d)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		CreatedAt:   createdAt,
		TotalAmount: totalAmount,
		Details:     insertedDetails,
	}, nil
}
