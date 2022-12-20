package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"strings"
)

type Transaction struct {
	senderAddress   string
	receiverAddress string
	amount          float32
}

func NewTransaction(sender string, receiver string, amount float32) *Transaction {
	return &Transaction{sender, receiver, amount}
}

func (t *Transaction) Print() {

	fmt.Printf("sender: %s\n", t.senderAddress)
	fmt.Printf("receiver: %s\n", t.receiverAddress)
	fmt.Printf("amount: %.3f\n", t.amount)
	println(color.Ize(color.Yellow, strings.Repeat("-", 22)))
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress   string  `json:"sender_address"`
		ReceiverAddress string  `json:"receiver_address"`
		Amount          float32 `json:"amount"`
	}{
		SenderAddress:   t.senderAddress,
		ReceiverAddress: t.receiverAddress,
		Amount:          t.amount,
	})
}

func (bc *Blockchain) CalculateBalance(blockchainAddress string) float32 {
	var balance float32 = 0.00
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			amount := t.amount
			if blockchainAddress == t.receiverAddress {
				balance += amount
			}

			if blockchainAddress == t.senderAddress {
				balance -= amount
			}
		}
	}

	return balance
}
