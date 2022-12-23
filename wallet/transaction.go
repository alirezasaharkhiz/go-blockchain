package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"go-blockchain/util"
)

type Transaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	senderAddress    string
	receiverAddress  string
	amount           float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, senderAddress string, receiverAddress string, amount float32) *Transaction {
	return &Transaction{
		amount:           amount,
		receiverAddress:  receiverAddress,
		senderAddress:    senderAddress,
		senderPrivateKey: privateKey,
		senderPublicKey:  publicKey,
	}
}

func (t *Transaction) SignTransaction() *util.Signature {
	mjt, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(mjt))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &util.Signature{R: r, S: s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender   string  `json:"sender_address"`
		Receiver string  `json:"receiver_address"`
		Amount   float32 `json:"amount"`
	}{
		Sender:   t.senderAddress,
		Receiver: t.receiverAddress,
		Amount:   t.amount,
	})
}
