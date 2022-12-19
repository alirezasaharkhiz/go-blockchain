package main

import (
	"github.com/TwiN/go-color"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := Block{}
	bc := Blockchain{}
	bc.AddBlock(0, b.Hash256())

	return &bc
}

func (bc *Blockchain) AddBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}

	return b
}

func (bc *Blockchain) AddTransaction(sender string, receiver string, amount float32) {
	t := NewTransaction(sender, receiver, amount)
	bc.transactionPool = append(bc.transactionPool, t)
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
