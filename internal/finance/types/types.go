package types

import (
	"github.com/pkg/errors"
)

type OperationType int

const (
	OperationTypeUndefined OperationType = iota
	OperationTypeReplenishment
	OperationTypeTransfer
)

func (t *OperationType) ToInt() int {
	return int(*t)
}

func NewOperationType(opType int) (OperationType, error) {
	switch opType {
	case 1:
		return OperationTypeReplenishment, nil
	case 2:
		return OperationTypeTransfer, nil
	default:
		return OperationTypeUndefined, errors.New("unknown operation type")
	}
}

type Balance int
