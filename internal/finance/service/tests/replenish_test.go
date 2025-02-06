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

func TestFinanceService_Replenish(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.Repository)
	mockTx := new(pkgMocks.TxManager)
	svc := service.NewFinanceService(mockRepo, mockTx)

	successID := uuid.New()

	testCases := []struct {
		name        string
		funds       *models.Funds
		mockSetup   func()
		expectedID  uuid.UUID
		expectedErr error
	}{
		{
			name: "success",
			funds: &models.Funds{
				UserID: uuid.New(),
				Amount: 100,
			},
			mockSetup: func() {
				mockRepo.On("Replenish", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("CreateOperation", mock.Anything, mock.Anything, types.OperationTypeReplenishment).
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
			name: "invalid userID error",
			funds: &models.Funds{
				UserID: uuid.Nil,
				Amount: 500,
			},
			mockSetup:   func() {},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrUserIDCanNotBeNil,
		},
		{
			name: "invalid amount error",
			funds: &models.Funds{
				UserID: successID,
				Amount: -500,
			},
			mockSetup:   func() {},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrAmountMustBePositive,
		},
		{
			name: "replenish error",
			funds: &models.Funds{
				UserID: uuid.New(),
				Amount: 100,
			},
			mockSetup: func() {
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
			expectedErr: errors.ErrReplenishFailed,
		},
		{
			name: "create operation error",
			funds: &models.Funds{
				UserID: uuid.New(),
				Amount: 100,
			},
			mockSetup: func() {
				mockRepo.On("Replenish", mock.Anything, mock.Anything).
					Return(nil)

				mockRepo.On("CreateOperation", mock.Anything, mock.Anything, types.OperationTypeReplenishment).
					Return(uuid.Nil, errors.ErrDatabaseError)

				mockTx.On("ReadCommitted", mock.Anything, mock.Anything).
					Return(errors.ErrDatabaseError).
					Run(func(args mock.Arguments) {
						fn := args.Get(1).(db.Handler)
						_ = fn(ctx)
					})
			},
			expectedID:  uuid.Nil,
			expectedErr: errors.ErrReplenishFailed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			mockTx.ExpectedCalls = nil

			tc.mockSetup()

			id, err := svc.Replenish(ctx, tc.funds)

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
