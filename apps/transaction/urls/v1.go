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

/*
https://www.sarulabs.com/post/5/2019-08-12/sending-docker-logs-to-elasticsearch-and-kibana-with-filebeat.html
*/
