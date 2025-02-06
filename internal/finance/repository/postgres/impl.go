package postgres

import (
	"context"
	"finance-operations-service/internal/finance/models"
	"finance-operations-service/internal/finance/types"
	"finance-operations-service/pkg/client/db"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (r *repo) Replenish(ctx context.Context, add *models.Funds) error {
	queryRaw := `UPDATE accounts SET balance = balance + $1 WHERE user_id = $2`

	q := db.Query{
		Name:     "operation_repository.Replenish",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, q, add.Amount, add.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CreateOperation(ctx context.Context, details *models.DetailsOperation, opType types.OperationType) (uuid.UUID, error) {
	queryRaw := `INSERT INTO operations (source_user_id, destination_user_id, type, amount)
					VALUES ($1, $2, $3, $4) RETURNING id`

	q := db.Query{
		Name:     "operation_repository.CreateOperation",
		QueryRaw: queryRaw,
	}

	var id uuid.UUID
	err := r.db.DB().QueryRowContext(ctx, q, details.SourceUserID,
		details.DestinationUserID, opType.ToInt(), details.Amount).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *repo) LastOperations(ctx context.Context, userID uuid.UUID) (models.Operations, error) {
	queryRow := `SELECT * FROM operations WHERE source_user_id = $1 OR destination_user_id = $1
                         ORDER BY request_time DESC LIMIT 10`

	q := db.Query{
		Name:     "operation_repository.LastOperations",
		QueryRaw: queryRow,
	}
	rows, err := r.db.DB().QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	ops := make(models.Operations, 0, 10)

	var (
		id        uuid.UUID
		srcUserID uuid.UUID
		dstUserID *uuid.UUID
		opType    int
		amount    int
		reqTime   time.Time
	)

	for rows.Next() {
		err = rows.Scan(&id, &srcUserID, &dstUserID, &opType, &amount, &reqTime)
		if err != nil {
			slog.Error("last_operations: failed to scan rows", "err", err)
			continue
		}

		operationType, err := types.NewOperationType(opType)
		if err != nil {
			slog.Error("last_operations: failed to get operation type", "err", err)
			continue
		}

		ops = append(ops, models.Operation{
			ID:          id,
			RequestTime: reqTime,
			Type:        operationType,
			Details: models.DetailsOperation{
				SourceUserID:      srcUserID,
				DestinationUserID: dstUserID,
				Amount:            amount,
			},
		})
	}

	return ops, nil
}

func (r *repo) GetAccountByUserID(ctx context.Context, userID uuid.UUID) (*models.Account, error) {
	queryRaw := `SELECT balance FROM accounts WHERE user_id = $1`

	q := db.Query{
		Name:     "operation_repository.GetAccountByUserID",
		QueryRaw: queryRaw,
	}

	var balance int64
	err := r.db.DB().QueryRowContext(ctx, q, userID).Scan(&balance)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		UserID:  userID,
		Balance: types.Balance(balance),
	}, nil
}

func (r *repo) Decrease(ctx context.Context, funds *models.Funds) error {
	queryRaw := `UPDATE accounts SET balance = balance - $1 WHERE user_id = $2`

	q := db.Query{
		Name:     "operation_repository.Decrease",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, q, funds.Amount, funds.UserID)
	if err != nil {
		return err
	}
	return nil
}
