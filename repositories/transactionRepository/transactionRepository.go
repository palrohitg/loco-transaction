package transactionRepository

import (
	"gorm.io/gorm"
	"loco-transaction/config"
	"loco-transaction/config/constants"
	"loco-transaction/dtos"
	"loco-transaction/models"
)

func CreateOrUpdateTransaction(db *gorm.DB, transaction dtos.Transaction) (models.TransactionSummary, error) {
	var transactionSummary models.TransactionSummary
	res := db.Table(constants.TransactionSummary).
		Where("transaction_id = ? ", transaction.Id).
		Find(&transactionSummary)
	if res.Error != nil {
		return transactionSummary, res.Error
	}
	if res.RowsAffected == 0 {
		return CreateTransaction(db, transaction)
	}
	if err := UpdateTransaction(db, &transactionSummary, transaction); err != nil {
		return transactionSummary, err
	}
	return transactionSummary, nil
}

func CreateTransaction(db *gorm.DB, transaction dtos.Transaction) (models.TransactionSummary, error) {
	transactionSummary := models.TransactionSummary{
		ParentId:      transaction.ParentId,
		TransactionId: transaction.Id,
		Amount:        transaction.Amount,
		Types:         transaction.Type,
	}
	if err := db.Table(constants.TransactionSummary).Create(&transactionSummary).Error; err != nil {
		return transactionSummary, err
	}
	return transactionSummary, nil
}

func UpdateTransaction(db *gorm.DB, transactionSummary *models.TransactionSummary, transaction dtos.Transaction) error {
	transactionSummary.Amount = transaction.Amount
	transactionSummary.Types = transaction.Type
	if err := db.Save(&transactionSummary).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactionByCategory(db *gorm.DB, category string) ([]models.TransactionSummary, error) {
	var transaction []models.TransactionSummary
	res := db.Table(constants.TransactionSummary).
		Where("types = ?", category).Find(&transaction)
	if res.Error != nil {
		return transaction, res.Error
	}
	if res.RowsAffected == 0 {
		return transaction, config.ValidationError(config.TransactionDoesNotExists)
	}
	return transaction, nil
}

func CalculateSumByTransaction(db *gorm.DB, transactionId int) (float64, error) {
	var sum float64
	tx := db.Begin()
	if err := db.Raw(`
         WITH RECURSIVE TransactionHierarchy AS (
             SELECT transaction_id, amount, parent_id
             FROM transaction_summary
             WHERE transaction_id = ?
             UNION ALL
             SELECT t.transaction_id, t.amount, t.parent_id
             FROM transaction_summary t
             INNER JOIN TransactionHierarchy th ON t.parent_id = th.transaction_id
         )
         SELECT SUM(amount) AS sum
         FROM TransactionHierarchy;
     `, transactionId).Scan(&sum).Error; err != nil {
		tx.Rollback()
		return 0, config.ValidationErrorReal(config.SumAPIError)
	}
	tx.Commit()
	return sum, nil
}
