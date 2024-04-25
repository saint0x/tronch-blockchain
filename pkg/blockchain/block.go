package blockchain

import (
  "crypto/sha256"
  "encoding/hex"
  "strconv"
  "strings"
  "time"
)

type Block struct {
  Index         int
  Timestamp     int64
  Transactions  []Transaction
  PrevBlockHash string
  Hash          string
  Nonce         int
  Difficulty    int
}


func CalculateHash(block Block) string {
  transactionStrings := block.TransactionsToStrings()
  record := strconv.Itoa(block.Index) + strconv.FormatInt(block.Timestamp, 10) + block.PrevBlockHash + strings.Join(transactionStrings, "") + strconv.Itoa(block.Nonce)
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}

func (block *Block) TransactionsToStrings() []string {
  var transactionStrings []string
  for _, tx := range block.Transactions {
    txString := tx.ID + tx.From + tx.To + strconv.FormatFloat(tx.Amount, 'f', -1, 64) + tx.Requirements + tx.Description
    transactionStrings = append(transactionStrings, txString)
  }
  return transactionStrings
}

func MineBlock(prevBlock Block, transactions []Transaction, difficulty int) Block {
  var newBlock Block
  newBlock.Index = prevBlock.Index + 1
  newBlock.Timestamp = time.Now().Unix()
  newBlock.Transactions = transactions
  newBlock.PrevBlockHash = prevBlock.Hash
  newBlock.Difficulty = difficulty
  newBlock.Nonce = 0

  for {
    hash := CalculateHash(newBlock)
    if IsValidHash(hash, difficulty) {
      newBlock.Hash = hash
      break
    }
    newBlock.Nonce++
  }

  return newBlock
}

func IsValidHash(hash string, difficulty int) bool {
  prefix := strings.Repeat("0", difficulty)
  return strings.HasPrefix(hash, prefix)
}
