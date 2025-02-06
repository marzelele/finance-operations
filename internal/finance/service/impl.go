package service

import (
	"context"
	"finance-operations-service/internal/finance/errors"
	"finance-operations-service/internal/finance/models"
	"finance-operations-service/internal/finance/types"
	"github.com/google/uuid"
	"log/slog"
)

func (s *service) Replenish(ctx context.Context, funds *models.Funds) (uuid.UUID, error) {
	err := funds.Validate()
	if err != nil {
		return uuid.Nil, err
	}

	var id uuid.UUID
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.financeRepository.Replenish(ctx, funds)
		if errTx != nil {
			slog.Error("replenish: failed to replenish funds", "err", errTx)
			return errTx
		}

		details := &models.DetailsOperation{
			SourceUserID:      funds.UserID,
			DestinationUserID: nil,
			Amount:            funds.Amount,
		}

		id, errTx = s.financeRepository.CreateOperation(ctx, details, types.OperationTypeReplenishment)
		if errTx != nil {
			slog.Error("replenish: failed to create operation", "err", errTx)
			return errTx
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, errors.ErrReplenishFailed
	}

	return id, nil
}

func (s *service) Transfer(ctx context.Context, details *models.DetailsOperation) (uuid.UUID, error) {
	err := details.Validate()
	if err != nil {
		return uuid.Nil, err
	}

	_, err = s.financeRepository.GetAccountByUserID(ctx, *details.DestinationUserID)
	if err != nil {
		slog.Error("destination user not found", "err", err)
		return uuid.Nil, errors.ErrDestinationUserNotFound
	}

	srcUser, err := s.financeRepository.GetAccountByUserID(ctx, details.SourceUserID)
	if err != nil {
		slog.Error("failed to get source user", "err", err)
		return uuid.Nil, errors.ErrSourceUserNotFound
	}

	if int(srcUser.Balance) < details.Amount {
		return uuid.Nil, errors.ErrNotEnoughFunds
	}

	var id uuid.UUID
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		decreaseFunds := &models.Funds{
			UserID: details.SourceUserID,
			Amount: details.Amount,
		}

		replenishFunds := &models.Funds{
			UserID: *details.DestinationUserID,
			Amount: details.Amount,
		}

		errTx := s.financeRepository.Replenish(ctx, replenishFunds)
		if errTx != nil {
			slog.Error("transfer: failed to replenish funds", "err", errTx)
			return errTx
		}

		errTx = s.financeRepository.Decrease(ctx, decreaseFunds)
		if errTx != nil {
			slog.Error("transfer: failed to decrease funds", "err", errTx)
			return errTx
		}

		id, errTx = s.financeRepository.CreateOperation(ctx, details, types.OperationTypeTransfer)
		if errTx != nil {
			slog.Error("transfer: failed to create operation", "err", errTx)
			return errTx
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, errors.ErrTransferFailed
	}

	return id, nil
}

func (s *service) LastOperations(ctx context.Context, userID uuid.UUID) (models.Operations, error) {
	ops, err := s.financeRepository.LastOperations(ctx, userID)
	if err != nil {
		slog.Error("failed to get last operations", "err", err)
		return models.Operations{}, errors.ErrFailedGetLastOperations
	}

	return ops, nil
}
