package transactionSerializer

import "loco-transaction/models"

func CreateOrUpdateSerializer(transactionSummary models.TransactionSummary) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	response["transaction_id"] = transactionSummary.TransactionId
	response["parent_id"] = transactionSummary.ParentId
	response["amount"] = transactionSummary.Amount
	response["parent_id"] = transactionSummary.ParentId
	response["type"] = transactionSummary.Types
	return response, nil
}

func TransactionSerializer(transactionSummary []models.TransactionSummary) ([]map[string]interface{}, error) {
	var response []map[string]interface{}
	for _, transaction := range transactionSummary {
		currentTransaction := make(map[string]interface{})
		currentTransaction["transaction_id"] = transaction.TransactionId
		currentTransaction["parent_id"] = transaction.ParentId
		currentTransaction["amount"] = transaction.Amount
		currentTransaction["parent_id"] = transaction.ParentId
		currentTransaction["type"] = transaction.Types
		response = append(response, currentTransaction)
	}
	return response, nil
}
