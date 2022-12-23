package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"github.com/TwiN/go-color"
	"go-blockchain/util"
	"log"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
	address         string
}

func NewBlockchain(address string) *Blockchain {
	b := Block{}
	bc := Blockchain{address: address}
	bc.AddBlock(0, b.Hash256())

	return &bc
}

func (bc *Blockchain) AddBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}

	return b
}

func (bc *Blockchain) AddTransaction(sender string, receiver string, amount float32, senderPublicKey *ecdsa.PublicKey, s *util.Signature) bool {
	t := NewTransaction(sender, receiver, amount)

	if sender == MiningSender {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		//
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("ERR : Unverified transaction ", t)
	}

	return false
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *util.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) CloneTransactionPool() []*Transaction {
	ts := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		ts = append(ts, NewTransaction(t.senderAddress, t.receiverAddress, t.amount))
	}

	return ts
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		println(color.Ize(color.Underline, strings.Repeat("-", 55)))
		println(color.Ize(color.Green, "Block"), i)
		block.Print()
	}

	println(color.Ize(color.Underline, strings.Repeat("#", 55)))
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
