package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"simple-cashier-api/models"
	"simple-cashier-api/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) HandleGetTodaysReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTodaysReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) GetTodaysReport(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now()

	report, err := h.service.GetTransactionReport(&timeNow, &timeNow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *TransactionHandler) HandleGetRangeDateTransactionReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetRangeDateTransactionReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) GetRangeDateTransactionReport(w http.ResponseWriter, r *http.Request) {
	startDateParam := r.URL.Query().Get("start_date")
	endDateParam := r.URL.Query().Get("end_date")

	var startDate, endDate *time.Time
	if startDateParam != "" {
		parsed, err := time.Parse("2006-01-02", startDateParam)
		if err != nil {
			http.Error(w, "Invalid start_date", http.StatusBadRequest)
			return
		}
		startDate = &parsed
	}
	if endDateParam != "" {
		parsed, err := time.Parse("2006-01-02", endDateParam)
		if err != nil {
			http.Error(w, "Invalid end_date", http.StatusBadRequest)
			return
		}
		endDate = &parsed
	}

	report, err := h.service.GetTransactionReport(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
