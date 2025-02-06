package postgres

import (
	"finance-operations-service/internal/finance"
	"finance-operations-service/pkg/client/db"
)

type repo struct {
	db db.Client
}

func NewFinanceRepository(db db.Client) finance.Repository {
	return &repo{db: db}
}
