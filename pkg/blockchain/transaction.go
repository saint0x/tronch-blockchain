// pkg/blockchain/transaction.go

package blockchain

import (
  "github.com/google/uuid"
)

type Transaction struct {
  ID           string
  From         string
  To           string
  Amount       float64
  Requirements string
  Description  string
  Signature    string
}

func NewTransaction(from, to string, amount float64, requirements, description string) Transaction {
  return Transaction{
    ID:           uuid.New().String(),
    From:         from,
    To:           to,
    Amount:       amount,
    Requirements: requirements,
    Description:  description,
    Signature:    "", // We'll implement signature later
  }
}
