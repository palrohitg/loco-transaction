package transactionController

import (
	"loco-transaction/apps/transaction/serializer/transactionSerializer"
	"loco-transaction/dbManager"
	"loco-transaction/repositories/transactionRepository"
)

func GetTransactionByType(category string) ([]map[string]interface{}, error) {
	var response []map[string]interface{}
	db, _ := dbManager.GetDbInstance()
	transactionSummary, err := transactionRepository.GetTransactionByCategory(db, category)
	if err != nil {
		return response, err
	}
	return transactionSerializer.TransactionSerializer(transactionSummary)
}
