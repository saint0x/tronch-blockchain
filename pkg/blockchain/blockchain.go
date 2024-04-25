// pkg/blockchain/blockchain.go

package blockchain

import (
    "time"
)

type Blockchain struct {
    Blocks []Block
}

func NewBlockchain() Blockchain {
    return Blockchain{
        Blocks: []Block{GenesisBlock()}, // We'll define the Genesis block later
    }
}

func (bc *Blockchain) AddBlock(transactions []Transaction, difficulty int) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := MineBlock(prevBlock, transactions, difficulty)
    bc.Blocks = append(bc.Blocks, newBlock)
}

func GenesisBlock() Block {
    return Block{
        Index:         0,
        Timestamp:     time.Now().Unix(),
        Transactions:  []Transaction{},
        PrevBlockHash: "",
        Hash:          "",
        Nonce:         0,
        Difficulty:    0,
    }
}

func (bc *Blockchain) IsValid() bool {
    for i := 1; i < len(bc.Blocks); i++ {
        currentBlock := bc.Blocks[i]
        prevBlock := bc.Blocks[i-1]

        if currentBlock.Hash != CalculateHash(currentBlock) {
            return false
        }

        if currentBlock.PrevBlockHash != prevBlock.Hash {
            return false
        }

        if !IsValidHash(currentBlock.Hash, currentBlock.Difficulty) {
            return false
        }
    }
    return true
}
