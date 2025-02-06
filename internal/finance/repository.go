package finance

import (
	"context"
	"finance-operations-service/internal/finance/models"
	"finance-operations-service/internal/finance/types"
	"github.com/google/uuid"
)

//go:generate mockery --name=Repository --filename=repository_mock.go --with-expecter

type Repository interface {
	Replenish(context.Context, *models.Funds) error
	CreateOperation(context.Context, *models.DetailsOperation, types.OperationType) (uuid.UUID, error)
	LastOperations(context.Context, uuid.UUID) (models.Operations, error)
	GetAccountByUserID(context.Context, uuid.UUID) (*models.Account, error)
	Decrease(context.Context, *models.Funds) error
}
