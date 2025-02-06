package service

import (
	"finance-operations-service/internal/finance"
	"finance-operations-service/pkg/client/db"
)

type service struct {
	financeRepository finance.Repository
	txManager         db.TxManager
}

func NewFinanceService(financeRepository finance.Repository, txManager db.TxManager) finance.Service {
	return &service{
		financeRepository: financeRepository,
		txManager:         txManager,
	}
}
