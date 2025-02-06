package tests

import (
	"context"
	"finance-operations-service/internal/finance/errors"
	"finance-operations-service/internal/finance/mocks"
	"finance-operations-service/internal/finance/models"
	"finance-operations-service/internal/finance/service"
	pkgMocks "finance-operations-service/pkg/client/db/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestFinanceService_LastOperations(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.Repository)
	mockTx := new(pkgMocks.TxManager)
	svc := service.NewFinanceService(mockRepo, mockTx)

	userID := uuid.New()

	expectedOperations := models.Operations{
		{
			ID:          uuid.New(),
			RequestTime: time.Now(),
			Details: models.DetailsOperation{
				SourceUserID:      uuid.New(),
				DestinationUserID: &userID,
				Amount:            100,
			},
		},
		{
			ID:          uuid.New(),
			RequestTime: time.Now(),
			Details: models.DetailsOperation{
				SourceUserID:      userID,
				DestinationUserID: &userID,
				Amount:            500,
			},
		},
	}

	testCases := []struct {
		name        string
		userID      uuid.UUID
		mockSetup   func()
		expectedOps models.Operations
		expectedErr error
	}{
		{
			name:   "success",
			userID: userID,
			mockSetup: func() {
				mockRepo.On("LastOperations", mock.Anything, userID).
					Return(expectedOperations, nil)
			},
			expectedOps: expectedOperations,
			expectedErr: nil,
		},
		{
			name:   "database error",
			userID: userID,
			mockSetup: func() {
				mockRepo.On("LastOperations", mock.Anything, userID).
					Return(models.Operations{}, errors.ErrDatabaseError)
			},
			expectedOps: models.Operations{},
			expectedErr: errors.ErrFailedGetLastOperations,
		},
		{
			name:   "empty operations",
			userID: uuid.Nil,
			mockSetup: func() {
				mockRepo.On("LastOperations", mock.Anything, uuid.Nil).
					Return(models.Operations{}, nil)
			},
			expectedOps: models.Operations{},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			mockTx.ExpectedCalls = nil

			tc.mockSetup()

			ops, err := svc.LastOperations(ctx, tc.userID)

			assert.Equal(t, tc.expectedOps, ops)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
