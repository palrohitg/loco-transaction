package transactionController

import (
	"loco-transaction/apps/transaction/serializer/transactionSerializer"
	"loco-transaction/dbManager"
	"loco-transaction/dtos"
	"loco-transaction/repositories/transactionRepository"
)

func UpdateOrCreateTransactionById(transaction dtos.Transaction) (map[string]interface{}, error) {
	var response map[string]interface{}
	db, _ := dbManager.GetDbInstance()
	transactionSummary, err := transactionRepository.CreateOrUpdateTransaction(db, transaction)
	if err != nil {
		return response, err
	}
	return transactionSerializer.CreateOrUpdateSerializer(transactionSummary)
}
