package finance

import (
	"context"
	"finance-operations-service/internal/finance/models"
	"github.com/google/uuid"
)

type Service interface {
	Replenish(context.Context, *models.Funds) (uuid.UUID, error)
	Transfer(ctx context.Context, details *models.DetailsOperation) (uuid.UUID, error)
	LastOperations(context.Context, uuid.UUID) (models.Operations, error)
}
