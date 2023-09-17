package urls

import (
	"loco-transaction/apps/transaction/handlers/transactionHandler"
	"loco-transaction/dtos"
)

func AddV1URLs(groups dtos.RouterGroups) {
	noAuthRg := groups.NoAuth
	TransactionGroup := noAuthRg.Group("/transaction-service")
	TransactionGroup.PUT("/transaction", transactionHandler.UpdateTransactionHandler)
	TransactionGroup.GET("/types/:types", transactionHandler.GetTransactionByCategoryHandler)
	TransactionGroup.GET("/sum/:transaction_id", transactionHandler.GetTransactionSumHandler)
}
