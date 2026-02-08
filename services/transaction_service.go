package services

import (
	"time"

	"simple-cashier-api/models"
	"simple-cashier-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetTransactionReport(startDate *time.Time, endDate *time.Time) (*models.TransactionReport, error) {
	sDate := time.Time{}
	eDate := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)

	if startDate != nil {
		sDate = *startDate
	}
	if endDate != nil {
		eDate = *endDate
	}

	return s.repo.GetTransactionReport(sDate, eDate)
}
