package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "tronch-blockchain/pkg/blockchain"
)

func main() {
  // Create a new blockchain with a genesis block
  bc := blockchain.NewBlockchain()

  // Test adding blocks
  fmt.Println("Adding Genesis block...")
  bc.AddBlock([]blockchain.Transaction{}, 2)
  if bc.IsValid() {
    fmt.Println("Genesis block added and blockchain is valid!")
  } else {
    fmt.Println("Error: Genesis block not added or blockchain is invalid!")
    return
  }

  // Test adding more blocks
  fmt.Println("Adding more blocks...")
  tx1 := blockchain.NewTransaction("Alice", "Bob", 1.0, "Requirements for tx1", "Description for tx1")
  tx2 := blockchain.NewTransaction("Bob", "Charlie", 2.0, "Requirements for tx2", "Description for tx2")
  transactions := []blockchain.Transaction{tx1, tx2}

  bc.AddBlock(transactions, 2)
  if bc.IsValid() {
    fmt.Println("Blocks added and blockchain is valid!")
  } else {
    fmt.Println("Error: Blocks not added or blockchain is invalid!")
    return
  }

  // Display the blockchain
  fmt.Println("Blockchain:")
  for _, block := range bc.Blocks {
    fmt.Printf("Index: %d\n", block.Index)
    fmt.Printf("Timestamp: %d\n", block.Timestamp)
    fmt.Printf("Hash: %s\n", block.Hash)
    fmt.Printf("PrevBlockHash: %s\n", block.PrevBlockHash)
    fmt.Printf("Transactions: %v\n", block.Transactions)
    fmt.Printf("Nonce: %d\n", block.Nonce)
    fmt.Printf("Difficulty: %d\n", block.Difficulty)
    fmt.Println("-------------------------------")
  }

  // Scanner to show blocks
  fmt.Println("(I made a blockchain today)")
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    input := strings.TrimSpace(scanner.Text())
    if input == "exit" {
      fmt.Println("Exiting...")
      break
    }

    // Refresh blockchain display
    fmt.Println("Blockchain:")
    for _, block := range bc.Blocks {
      fmt.Printf("Index: %d\n", block.Index)
      fmt.Printf("Timestamp: %d\n", block.Timestamp)
      fmt.Printf("Hash: %s\n", block.Hash)
      fmt.Printf("PrevBlockHash: %s\n", block.PrevBlockHash)
      fmt.Printf("Transactions: %v\n", block.Transactions)
      fmt.Printf("Nonce: %d\n", block.Nonce)
      fmt.Printf("Difficulty: %d\n", block.Difficulty)
      fmt.Println("-------------------------------")
    }
    fmt.Println("I made a blockchain today!")
  }
}
