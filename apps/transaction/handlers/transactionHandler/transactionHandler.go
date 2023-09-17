package transactionHandler

import (
	"github.com/gin-gonic/gin"
	"loco-transaction/apps/transaction/controller/transactionController"
	"loco-transaction/config"
	"loco-transaction/dtos"
	"net/http"
	"strconv"
)

func UpdateTransactionHandler(ctx *gin.Context) {
	var transaction dtos.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	data, err := transactionController.UpdateOrCreateTransactionById(transaction)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	config.SendSuccessResponse(ctx, config.SuccessStatus, config.SuccessMsg, data)
}

func GetTransactionByCategoryHandler(ctx *gin.Context) {
	data, err := transactionController.GetTransactionByType(ctx.Param("types"))
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	response := map[string]interface{}{
		"status":  config.SuccessStatus,
		"message": config.SuccessMsg,
		"data":    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func GetTransactionSumHandler(ctx *gin.Context) {
	transactionId, _ := strconv.Atoi(ctx.Param("transaction_id"))
	data, err := transactionController.TransactionSum(transactionId)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	response := map[string]interface{}{
		"status":  config.SuccessStatus,
		"message": config.SuccessMsg,
		"sum":     data,
	}
	ctx.JSON(http.StatusOK, response)
}
