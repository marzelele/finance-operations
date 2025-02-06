package errors

import "errors"

var (
	ErrAmountMustBePositive         = errors.New("amount must be positive")
	ErrSourceUserNotFound           = errors.New("source user not found")
	ErrDestinationUserNotFound      = errors.New("destination user not found")
	ErrReplenishFailed              = errors.New("replenish failed")
	ErrTransferFailed               = errors.New("transfer failed")
	ErrFailedGetLastOperations      = errors.New("failed to get last operations")
	ErrNotEnoughFunds               = errors.New("not enough funds")
	ErrUserIDCanNotBeNil            = errors.New("user id can not be nil")
	ErrSourceUserIDCanNotBeNil      = errors.New("source user id can not be nil")
	ErrDestinationUserIDCanNotBeNil = errors.New("destination user id can not be nil")
	ErrUsersAreSame                 = errors.New("source_user_id and destination_user_id cannot be the same")
	ErrDatabaseError                = errors.New("database error")
)
