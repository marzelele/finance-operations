package handlers

import (
	"finance-operations-service/internal/finance"
)

type Handler struct {
	financeService finance.Service
}

func NewHandler(service finance.Service) *Handler {
	return &Handler{
		financeService: service,
	}
}
