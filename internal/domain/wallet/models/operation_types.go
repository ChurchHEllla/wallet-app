package models

type OperationType string

const (
	OperationType_Deposit  OperationType = "DEPOSIT"
	OperationType_Withdraw OperationType = "WITHDRAW"
)

var OperationType_ALL = map[OperationType]struct{}{
	OperationType_Deposit:  {},
	OperationType_Withdraw: {},
}

func (v OperationType) Valid() bool {
	_, ok := OperationType_ALL[v]
	return ok
}
