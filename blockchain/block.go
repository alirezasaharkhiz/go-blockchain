package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		previousHash: previousHash,
		nonce:        nonce,
		transactions: transactions,
	}
}

func (b *Block) Hash256() [32]byte {
	mb, _ := json.Marshal(b)
	return sha256.Sum256([]byte(mb))
}

func (b *Block) Print() {
	fmt.Printf("previous hash: %x\n", b.previousHash)
	fmt.Printf("timestamp:     %d\n", b.timestamp)
	fmt.Printf("nonce:         %d\n", b.nonce)
	println(color.Ize(color.Yellow, "Transactions:"))
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PreviousHash [32]byte       `json:"previous_hash"`
		Nonce        int            `json:"nonce"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		PreviousHash: b.previousHash,
		Nonce:        b.nonce,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}
