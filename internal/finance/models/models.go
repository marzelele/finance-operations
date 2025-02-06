package models

import (
	"finance-operations-service/internal/finance/errors"
	"finance-operations-service/internal/finance/types"
	"github.com/google/uuid"
	"time"
)

type Funds struct {
	UserID uuid.UUID `json:"user_id"`
	Amount int       `json:"amount"`
}

func (a *Funds) Validate() error {
	if a.UserID == uuid.Nil {
		return errors.ErrUserIDCanNotBeNil
	}

	if a.Amount <= 0 {
		return errors.ErrAmountMustBePositive
	}

	return nil
}

type Operation struct {
	ID          uuid.UUID           `json:"id"`
	RequestTime time.Time           `json:"request_time"`
	Type        types.OperationType `json:"type"`
	Details     DetailsOperation    `json:"details"`
}

type Operations []Operation

type DetailsOperation struct {
	SourceUserID      uuid.UUID  `json:"source_user_id"`
	DestinationUserID *uuid.UUID `json:"destination_user_id"`
	Amount            int        `json:"amount"`
}

func (d *DetailsOperation) Validate() error {
	if d.SourceUserID == uuid.Nil {
		return errors.ErrSourceUserIDCanNotBeNil
	}

	if d.DestinationUserID == nil || *d.DestinationUserID == uuid.Nil {
		return errors.ErrDestinationUserIDCanNotBeNil
	}

	if d.SourceUserID == *d.DestinationUserID {
		return errors.ErrUsersAreSame
	}

	if d.Amount <= 0 {
		return errors.ErrAmountMustBePositive
	}

	return nil
}

type Account struct {
	UserID  uuid.UUID     `json:"id"`
	Balance types.Balance `json:"balance"`
}
