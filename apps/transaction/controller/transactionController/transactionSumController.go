package transactionController

import (
	"loco-transaction/dbManager"
	"loco-transaction/repositories/transactionRepository"
)

func TransactionSum(transactionId int) (float64, error) {
	db, _ := dbManager.GetDbInstance()
	sum, err := transactionRepository.CalculateSumByTransaction(db, transactionId)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
