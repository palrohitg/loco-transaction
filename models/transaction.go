package models

import (
	"time"
)

type TransactionSummary struct {
	Id            int       `json:"id"`
	TransactionId int       `json:"transaction_id"`
	ParentId      int       `json:"parent_id"`
	Types         string    `json:"types"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (b *TransactionSummary) TableName() string {
	return "transaction_summary"
}
