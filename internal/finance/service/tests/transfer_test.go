package tests

import (
	"context"
	"finance-operations-service/internal/finance/errors"
	"finance-operations-service/internal/finance/mocks"
	"finance-operations-service/internal/finance/models"
	"finance-operations-service/internal/finance/service"
	"finance-operations-service/internal/finance/types"
	"finance-operations-service/pkg/client/db"
	pkgMocks "finance-operations-service/pkg/client/db/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFinanceService_Transfer(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.Repository)
	mockTx := new(pkgMocks.TxManager)
	svc := service.NewFinanceService(mockRepo, mockTx)

	successID := uuid.New()
	sourceUserID := uuid.New()
	destinationUserID := uuid.New()

	testCases := []struct {
		name        string
		details     *models.DetailsOperation
		mockSetup   func()
		expectedID  uuid.UUID
		expectedErr error
	}{
		{
			name: "success",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            100,
			},
			mockSetup: func() {
				mockRepo.On("GetAccountByUserID", mock.Anything, destinationUserID).
					Return(&models.Account{
						UserID:  destinationUserID,
						Balance: types.Balance(500),
					}, nil)

				mockRepo.On("GetAccountByUserID", mock.Anything, sourceUserID).
					Return(&models.Account{
						UserID:  sourceUserID,
						Balance: types.Balance(200),
					}, nil)

				mockRepo.On("Replenish", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("Decrease", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("CreateOperation", mock.Anything, mock.Anything, types.OperationTypeTransfer).
					Return(successID, nil)

				mockTx.On("ReadCommitted", mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						fn := args.Get(1).(db.Handler)
						_ = fn(ctx)
					})
			},
			expectedID:  successID,
			expectedErr: nil,
		},
		{
			name: "invalid sourceUserID error",
			details: &models.DetailsOperation{
				SourceUserID:      uuid.Nil,
				DestinationUserID: &destinationUserID,
				Amount:            0,
			},
			mockSetup:   func() {},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrSourceUserIDCanNotBeNil,
		},
		{
			name: "invalid destinationUserID error",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: nil,
				Amount:            0,
			},
			mockSetup:   func() {},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrDestinationUserIDCanNotBeNil,
		},
		{
			name: "invalid amount error",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            -100,
			},
			mockSetup:   func() {},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrAmountMustBePositive,
		},
		{
			name: "destination user not found",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            100,
			},
			mockSetup: func() {
				mockRepo.On("GetAccountByUserID", mock.Anything, destinationUserID).
					Return(nil, errors.ErrDestinationUserNotFound)
			},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrDestinationUserNotFound,
		},
		{
			name: "not enough funds",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            300,
			},
			mockSetup: func() {
				mockRepo.On("GetAccountByUserID", mock.Anything, destinationUserID).
					Return(&models.Account{
						UserID:  destinationUserID,
						Balance: types.Balance(100),
					}, nil)

				mockRepo.On("GetAccountByUserID", mock.Anything, sourceUserID).
					Return(&models.Account{
						UserID:  sourceUserID,
						Balance: types.Balance(200),
					}, nil)
			},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrNotEnoughFunds,
		},
		{
			name: "replenish error",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            100,
			},
			mockSetup: func() {
				mockRepo.On("GetAccountByUserID", mock.Anything, destinationUserID).
					Return(&models.Account{
						UserID:  destinationUserID,
						Balance: types.Balance(500),
					}, nil)

				mockRepo.On("GetAccountByUserID", mock.Anything, sourceUserID).
					Return(&models.Account{
						UserID:  sourceUserID,
						Balance: types.Balance(200),
					}, nil)

				mockRepo.On("Replenish", mock.Anything, mock.Anything).
					Return(errors.ErrDatabaseError)

				mockTx.On("ReadCommitted", mock.Anything, mock.Anything).
					Return(errors.ErrDatabaseError).
					Run(func(args mock.Arguments) {
						fn := args.Get(1).(db.Handler)
						_ = fn(ctx)
					})
			},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrTransferFailed,
		},
		{
			name: "decrease error",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            100,
			},
			mockSetup: func() {
				mockRepo.On("GetAccountByUserID", mock.Anything, destinationUserID).
					Return(&models.Account{
						UserID:  destinationUserID,
						Balance: types.Balance(500),
					}, nil)

				mockRepo.On("GetAccountByUserID", mock.Anything, sourceUserID).
					Return(&models.Account{
						UserID:  sourceUserID,
						Balance: types.Balance(200),
					}, nil)

				mockRepo.On("Replenish", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("Decrease", mock.Anything, mock.Anything).
					Return(errors.ErrDatabaseError)

				mockTx.On("ReadCommitted", mock.Anything, mock.Anything).
					Return(errors.ErrDatabaseError).
					Run(func(args mock.Arguments) {
						fn := args.Get(1).(db.Handler)
						_ = fn(ctx)
					})
			},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrTransferFailed,
		},
		{
			name: "create operation error",
			details: &models.DetailsOperation{
				SourceUserID:      sourceUserID,
				DestinationUserID: &destinationUserID,
				Amount:            100,
			},
			mockSetup: func() {
				mockRepo.On("GetAccountByUserID", mock.Anything, destinationUserID).
					Return(&models.Account{
						UserID:  destinationUserID,
						Balance: types.Balance(500),
					}, nil)

				mockRepo.On("GetAccountByUserID", mock.Anything, sourceUserID).
					Return(&models.Account{
						UserID:  sourceUserID,
						Balance: types.Balance(200),
					}, nil)

				mockRepo.On("Replenish", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("Decrease", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("CreateOperation", mock.Anything, mock.Anything, types.OperationTypeTransfer).
					Return(uuid.Nil, errors.ErrDatabaseError)

				mockTx.On("ReadCommitted", mock.Anything, mock.Anything).
					Return(errors.ErrDatabaseError).
					Run(func(args mock.Arguments) {
						fn := args.Get(1).(db.Handler)
						_ = fn(ctx)
					})
			},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrTransferFailed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			mockTx.ExpectedCalls = nil

			tc.mockSetup()

			id, err := svc.Transfer(ctx, tc.details)

			assert.Equal(t, tc.expectedID, id)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockTx.AssertExpectations(t)
		})
	}
}
