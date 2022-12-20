package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchainAddress := "my_blockchain_address"
	bc := blockchain.NewBlockchain(blockchainAddress)

	bc.AddTransaction("A", "B", 1.0)
	bc.AddTransaction("x", "z", 5.0)
	bc.AddTransaction("x", "A", 7.1)
	bc.Mining()
	//lastBlockHash := bc.LastBlock().Hash256()
	//nonce := bc.GenerateProofOfWorkNonce()
	//bc.AddBlock(nonce, lastBlockHash)

	bc.AddTransaction("c", "g", 5.0)
	bc.AddTransaction("c", "f", 7.1)
	bc.Mining()
	//lastBlockHash = bc.LastBlock().Hash256()
	//nonce = bc.GenerateProofOfWorkNonce()
	//bc.AddBlock(nonce, lastBlockHash)

	bc.Print()

	fmt.Println("c", bc.CalculateBalance("c"))
	fmt.Println("f", bc.CalculateBalance("f"))
	fmt.Println("A", bc.CalculateBalance("A"))
	fmt.Println("x", bc.CalculateBalance("x"))
}
