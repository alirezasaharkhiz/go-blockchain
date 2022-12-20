package blockchain

import (
	"fmt"
	"strings"
)

const RequiredProofChar string = "0"

func (bc *Blockchain) GenerateProofOfWorkNonce() int {
	ts := bc.CloneTransactionPool()
	ph := bc.LastBlock().Hash256()

	nonce := 0
	for !bc.IsValidProof(nonce, ph, ts, MiningDifficulty, RequiredProofChar) {
		nonce += 1
	}

	return nonce
}

func (bc *Blockchain) IsValidProof(nonce int, previousBlockHash [32]byte, transactions []*Transaction, difficulty int, requiredChar string) bool {
	requiredProofString := strings.Repeat(requiredChar, difficulty)
	guessingBlock := Block{nonce: nonce, previousHash: previousBlockHash, transactions: transactions, timestamp: 0}
	guessingBlockHash := fmt.Sprintf("%x", guessingBlock.Hash256())

	return guessingBlockHash[:difficulty] == requiredProofString
}
